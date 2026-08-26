package gen

import (
	"errors"
	"fmt"
	"net"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

var databaseIdentifierPattern = regexp.MustCompile(`^[A-Za-z_][A-Za-z0-9_]*$`)
var databaseHostPattern = regexp.MustCompile(`^[A-Za-z0-9][A-Za-z0-9.-]*$`)

func databaseURLMissingError() error {
	return fmt.Errorf("database configuration error: DATABASE_URL is not set\n\nSet it to a valid connection URL, for example:\n  mysql://user:password@localhost:3306/database")
}

func invalidDatabaseURLError() error {
	return fmt.Errorf("database configuration error: DATABASE_URL is invalid\n\nExpected a URL such as:\n  mysql://user:password@localhost:3306/database")
}

func invalidDatabaseURLDetailError(detail string) error {
	return fmt.Errorf("database configuration error: DATABASE_URL %s\n\nExpected a URL such as:\n  mysql://user:password@localhost:3306/database", detail)
}

func unsupportedDatabaseSchemeError(scheme string) error {
	if strings.TrimSpace(scheme) == "" {
		scheme = "(empty)"
	}
	return fmt.Errorf("unsupported database type %q; supported types are mysql, postgres, sqlite3", scheme)
}

func validateDatabaseURL(rawURL string, databaseURL *url.URL) error {
	if databaseURL == nil {
		return invalidDatabaseURLError()
	}
	if !strings.Contains(rawURL, "://") {
		return invalidDatabaseURLDetailError("must include a database scheme such as mysql://")
	}

	switch databaseURL.Scheme {
	case "mysql", "postgres":
		if strings.TrimSpace(databaseURL.Hostname()) == "" {
			return invalidDatabaseURLDetailError(fmt.Sprintf("for %s requires a host", databaseURL.Scheme))
		}
		if port := databaseURL.Port(); port != "" && !validDatabasePort(port) {
			return invalidDatabaseURLDetailError(fmt.Sprintf("for %s has an invalid port", databaseURL.Scheme))
		}
		if strings.Trim(strings.TrimPrefix(databaseURL.Path, "/"), " ") == "" {
			return invalidDatabaseURLDetailError(fmt.Sprintf("for %s requires a database name", databaseURL.Scheme))
		}
	case "sqlite3":
		if strings.TrimSpace(databaseURL.Host+databaseURL.Path) == "" {
			return invalidDatabaseURLDetailError("for sqlite3 requires a database file")
		}
	default:
		return unsupportedDatabaseSchemeError(databaseURL.Scheme)
	}

	return nil
}

func formatDatabaseConnectionError(databaseURL *url.URL, connectionErr error) error {
	if connectionErr == nil {
		return nil
	}

	scheme, databaseName, host := "database", "", ""
	if databaseURL != nil {
		scheme = databaseURL.Scheme
		databaseName = strings.TrimPrefix(databaseURL.Path, "/")
		if decoded, err := url.PathUnescape(databaseName); err == nil {
			databaseName = decoded
		}
		host = databaseURL.Host
	}

	cause := redactDatabaseError(databaseURL, connectionErr.Error())
	lowerCause := strings.ToLower(cause)

	if databaseName != "" && (strings.Contains(lowerCause, "error 1049") || strings.Contains(lowerCause, "unknown database") || strings.Contains(lowerCause, "database does not exist")) {
		message := fmt.Sprintf("database %q does not exist", databaseName)
		if scheme == "mysql" {
			if command, ok := mysqlCreateDatabaseCommand(databaseURL, databaseName); ok {
				return fmt.Errorf("%s\n\nCreate it first:\n  %s\n\nThen initialize the schema:\n  make migrate\n\nAfter that, start the server again:\n  make start", message, command)
			}
		}
		return fmt.Errorf("%s\n\nCreate the database with your database administration tool, then run:\n  make migrate\n  make start", message)
	}

	if strings.Contains(lowerCause, "error 1045") || strings.Contains(lowerCause, "access denied") || strings.Contains(lowerCause, "authentication failed") || strings.Contains(lowerCause, "password authentication failed") {
		return fmt.Errorf("database authentication failed\n\nCheck the username and password in DATABASE_URL and verify that the database user has permission to connect")
	}

	if isDatabaseNetworkError(connectionErr, lowerCause) {
		if host == "" {
			host = "the configured host"
		}
		return fmt.Errorf("cannot reach the database server at %s\n\nCheck that the database service is running and that the host and port in DATABASE_URL are correct", host)
	}

	return fmt.Errorf("database connection failed (%s): %s\n\nCheck DATABASE_URL, database availability, and user permissions", scheme, cause)
}

func mysqlCreateDatabaseCommand(databaseURL *url.URL, databaseName string) (string, bool) {
	if databaseURL == nil || !databaseIdentifierPattern.MatchString(databaseName) {
		return "", false
	}

	host := databaseURL.Hostname()
	if !databaseHostPattern.MatchString(host) {
		return "", false
	}
	port := databaseURL.Port()
	if port == "" {
		port = "3306"
	}
	if !validDatabasePort(port) {
		return "", false
	}

	username := "root"
	if databaseURL.User != nil {
		configuredUsername := databaseURL.User.Username()
		if !databaseIdentifierPattern.MatchString(configuredUsername) {
			return "", false
		}
		username = configuredUsername
	}

	return fmt.Sprintf("mysql -h %s -P %s -u %s -p -e 'CREATE DATABASE %s CHARACTER SET utf8mb4'", host, port, username, databaseName), true
}

func validDatabasePort(port string) bool {
	portNumber, err := strconv.Atoi(port)
	return err == nil && portNumber >= 1 && portNumber <= 65535
}

func isDatabaseNetworkError(connectionErr error, lowerCause string) bool {
	var networkErr net.Error
	if errors.As(connectionErr, &networkErr) {
		return true
	}

	for _, fragment := range []string{
		"connection refused",
		"connection reset",
		"context deadline exceeded",
		"i/o timeout",
		"network is unreachable",
		"no such host",
		"operation not permitted",
		"operation timed out",
		"temporary failure in name resolution",
	} {
		if strings.Contains(lowerCause, fragment) {
			return true
		}
	}
	return false
}

func redactDatabaseError(databaseURL *url.URL, message string) string {
	if databaseURL == nil || databaseURL.User == nil {
		return message
	}
	password, hasPassword := databaseURL.User.Password()
	if !hasPassword || password == "" {
		return message
	}

	redacted := strings.ReplaceAll(message, password, "[REDACTED]")
	redacted = strings.ReplaceAll(redacted, url.QueryEscape(password), "[REDACTED]")
	redacted = strings.ReplaceAll(redacted, url.PathEscape(password), "[REDACTED]")
	return redacted
}
