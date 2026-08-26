package gen

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/cloudevents/sdk-go/pkg/cloudevents/transport"
	"github.com/gofrs/uuid"
	cloudeventsaws "github.com/maiguangyang/cloudevents-aws-transport"
)

const (
	ORMChangeEvent = "com.graphql.orm.change"
)

// EventType ...
type EventType string

const (
	// EventTypeCreated ...
	EventTypeCreated = "CREATED"
	// EventTypeUpdated ...
	EventTypeUpdated = "UPDATED"
	// EventTypeDeleted ...
	EventTypeDeleted = "DELETED"
)

type EventDataValue interface{}

// EventChange ...
type EventChange struct {
	Name     string "json:'name'"
	OldValue string "json:'oldValue'"
	NewValue string "json:'newValue'"
}

type EventController struct {
	clients map[string]cloudevents.Client
	debug   bool
	source  string
}

func NewEventController() (ec EventController, err error) {
	source := os.Getenv("EVENT_TRANSPORT_SOURCE")
	if source == "" {
		hostname, _err := os.Hostname()
		if err != nil {
			err = _err
			return
		}
		source = "http://" + hostname + "/graphql"
	}

	URLs := getENVArray("EVENT_TRANSPORT_URL")
	_clients := map[string]cloudevents.Client{}
	for _, URL := range URLs {
		if URL != "" {
			t, tErr := transportForURL(URL)
			err = tErr
			if err != nil {
				return
			}

			client, cErr := cloudevents.NewClient(t)
			err = cErr
			if err != nil {
				return
			}
			log.Printf("Created cloudevents client with target %s", URL)
			_clients[URL] = client
		}
	}
	debug := os.Getenv("DEBUG") == "true"
	ec = EventController{clients: _clients, debug: debug, source: source}

	log.Printf("Created EventController with source %s", source)

	return
}

func (c *EventController) send(ctx context.Context, e cloudevents.Event) error {
	for URL, client := range c.clients {
		if _, _, err := client.Send(ctx, e); err != nil {
			if c.debug {
				fmt.Printf("received cloudevents error %s from server %s\n", err.Error(), URL)
			}
			return err
		}
	}
	return nil
}

// SendEvent ...
func (c *EventController) SendEvent(ctx context.Context, e *Event) (err error) {
	if len(c.clients) == 0 {
		return
	}
	event := cloudevents.NewEvent()
	event.SetID(e.ID)
	event.SetType(ORMChangeEvent)
	event.SetSource(c.source)
	// event.SetTime(e.Date)
	err = event.SetData(e)
	if err != nil {
		return
	}
	err = c.send(ctx, event)
	return
}

func getENVArray(name string) []string {
	arr := []string{}

	val := os.Getenv(name)
	if val != "" {
		arr = append(arr, strings.Split(val, ",")...)
	}

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("%s_%d", name, i)
		sval := os.Getenv(key)
		if sval != "" {
			arr = append(arr, sval)
		}
	}

	return arr
}

func transportForURL(URL string) (t transport.Transport, err error) {

	if strings.HasPrefix(URL, "arn:aws:sns") {
		t, err = cloudeventsaws.NewSNSTransport(URL)
		return
	}
	if strings.HasPrefix(URL, "arn:aws:events") {
		t, err = cloudeventsaws.NewEventBridgeTransport(URL)
		return
	}

	u, err := url.Parse(URL)
	if err != nil {
		return
	}
	switch u.Scheme {
	case "http", "https":
		t, err = cloudevents.NewHTTPTransport(
			cloudevents.WithTarget(URL),
			cloudevents.WithBinaryEncoding(),
		)
	case "sqs+https":
		u.Scheme = "https"
		t, err = cloudeventsaws.NewSQSTransport(u.String())
	default:
		err = fmt.Errorf("unknown scheme %s", u.Scheme)

	}
	return
}

func (ec *EventChange) SetOldValue(value interface{}) error {
	data, err := json.Marshal(value)
	if err == nil {
		ec.OldValue = string(data)
	}
	return err
}
func (ec *EventChange) OldValueAs(data interface{}) error {
	return json.Unmarshal([]byte(ec.OldValue), data)
}
func (ec *EventChange) SetNewValue(value interface{}) error {
	data, err := json.Marshal(value)
	if err == nil {
		ec.NewValue = string(data)
	}
	return err
}
func (ec *EventChange) NewValueAs(data interface{}) error {
	return json.Unmarshal([]byte(ec.NewValue), data)
}

type EventMetadata struct {
	Type        EventType "json:'type'"
	Cursor      string    "json:'cursor'"
	Entity      string    "json:'entity'"
	EntityID    string    "json:'entityId'"
	Date        int64     "json:'date'"
	PrincipalID *string   "json:'principalId'"
}

// Event ...
type Event struct {
	EventMetadata
	ID      string         "json:'id'"
	Changes []*EventChange "json:'changes'"
}

// NewEvent ...
func NewEvent(meta EventMetadata) Event {
	return Event{
		EventMetadata: meta,
		ID:            uuid.Must(uuid.NewV4()).String(),
		Changes:       []*EventChange{},
	}
}

// HasChangedColumn check if given event has changes on specific column
func (e Event) HasChangedColumn(c string) bool {
	for _, col := range e.ChangedColumns() {
		if col == c {
			return true
		}
	}
	return false
}

// ChangedColumns returns list of names of changed columns
func (e Event) ChangedColumns() []string {
	columns := []string{}

	for _, change := range e.Changes {
		columns = append(columns, change.Name)
	}

	return columns
}

func (e *Event) Change(column string) (ec *EventChange) {
	for _, c := range e.Changes {
		if c.Name == column {
			ec = c
			break
		}
	}
	return
}

// AddNewValue ...
func (e *Event) AddNewValue(column string, v EventDataValue) {
	change := e.Change(column)
	if change == nil {
		c := EventChange{Name: column}
		change = &c
		e.Changes = append(e.Changes, change)
	}
	if err := change.SetNewValue(v); err != nil {
		panic("failed to set new value" + err.Error())
	}
}

// AddOldValue ...
func (e *Event) AddOldValue(column string, v EventDataValue) {
	change := e.Change(column)
	if change == nil {
		c := EventChange{Name: column}
		change = &c
		e.Changes = append(e.Changes, change)
	}
	if err := change.SetOldValue(v); err != nil {
		panic("failed to set new value" + err.Error())
	}
}
