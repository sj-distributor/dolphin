package templates

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestDatabaseErrorsTemplateReturnsFriendlyActionableMessages(t *testing.T) {
	program := strings.Replace(DatabaseErrors, "package gen", "package main", 1) + `

func requireContains(message, fragment string) {
	if !strings.Contains(message, fragment) {
		panic(fmt.Sprintf("message %q does not contain %q", message, fragment))
	}
}

func requireNotContains(message, fragment string) {
	if strings.Contains(message, fragment) {
		panic(fmt.Sprintf("message %q unexpectedly contains %q", message, fragment))
	}
}

func main() {
	missing := databaseURLMissingError().Error()
	requireContains(missing, "DATABASE_URL is not set")
	requireContains(missing, "mysql://user:password@localhost:3306/database")

	unsupported := unsupportedDatabaseSchemeError("oracle").Error()
	requireContains(unsupported, "oracle")
	requireContains(unsupported, "mysql, postgres, sqlite3")

	for raw, fragment := range map[string]string{
		"localhost:3306/dolphin_example": "must include a database scheme",
		"mysql://":                       "requires a host",
		"mysql://root@localhost:3306":    "requires a database name",
		"mysql://root@localhost:70000/db": "has an invalid port",
		"postgres://localhost":           "requires a database name",
		"sqlite3://":                     "requires a database file",
		"sqlite3://?cache=shared":         "requires a database file",
	} {
		parsed, _ := url.Parse(raw)
		validationErr := validateDatabaseURL(raw, parsed)
		if validationErr == nil {
			panic(fmt.Sprintf("validateDatabaseURL(%q) returned nil", raw))
		}
		requireContains(validationErr.Error(), fragment)
	}

	mysqlURL, _ := url.Parse("mysql://root:secret@localhost:3306/dolphin_example")
	missingDB := formatDatabaseConnectionError(mysqlURL, fmt.Errorf("Error 1049 (42000): Unknown database 'dolphin_example'")).Error()
	requireContains(missingDB, "database \"dolphin_example\" does not exist")
	requireContains(missingDB, "CREATE DATABASE dolphin_example CHARACTER SET utf8mb4")
	requireContains(missingDB, "-h localhost -P 3306")
	requireContains(missingDB, "make migrate")
	requireContains(missingDB, "make start")
	requireNotContains(missingDB, "secret")

	remoteURL, _ := url.Parse("mysql://app:secret@db.example.com:3307/dolphin_example")
	remoteMissingDB := formatDatabaseConnectionError(remoteURL, fmt.Errorf("Error 1049 (42000): Unknown database 'dolphin_example'")).Error()
	requireContains(remoteMissingDB, "mysql -h db.example.com -P 3307 -u app -p")

	unsafeHostURL, _ := url.Parse("mysql://root:secret@-option.example:3306/dolphin_example")
	unsafeHost := formatDatabaseConnectionError(unsafeHostURL, fmt.Errorf("Error 1049 (42000): Unknown database 'dolphin_example'")).Error()
	requireContains(unsafeHost, "database administration tool")
	requireNotContains(unsafeHost, "CREATE DATABASE")

	unsafeUserURL, _ := url.Parse("mysql://-option:secret@localhost:3306/dolphin_example")
	unsafeUser := formatDatabaseConnectionError(unsafeUserURL, fmt.Errorf("Error 1049 (42000): Unknown database 'dolphin_example'")).Error()
	requireContains(unsafeUser, "database administration tool")
	requireNotContains(unsafeUser, "CREATE DATABASE")

	auth := formatDatabaseConnectionError(mysqlURL, fmt.Errorf("Error 1045 (28000): Access denied for user 'root'@'localhost'")).Error()
	requireContains(auth, "database authentication failed")
	requireContains(auth, "username and password")
	requireNotContains(auth, "secret")

	refused := formatDatabaseConnectionError(mysqlURL, fmt.Errorf("dial tcp 127.0.0.1:3306: connect: connection refused")).Error()
	requireContains(refused, "cannot reach the database server")
	requireContains(refused, "localhost:3306")

	for _, networkErr := range []error{
		fmt.Errorf("dial tcp: i/o timeout"),
		fmt.Errorf("context deadline exceeded"),
		fmt.Errorf("read: connection reset by peer"),
		&net.DNSError{Err: "temporary failure in name resolution", Name: "db.example.com", IsTimeout: true},
	} {
		message := formatDatabaseConnectionError(remoteURL, networkErr).Error()
		requireContains(message, "cannot reach the database server")
	}

	generic := formatDatabaseConnectionError(mysqlURL, fmt.Errorf("connection failed with password secret")).Error()
	requireContains(generic, "database connection failed")
	requireContains(generic, "[REDACTED]")
	requireNotContains(generic, "secret")

	encodedPasswordURL, _ := url.Parse("mysql://root:p%40ss%3Aword@localhost:3306/dolphin_example")
	encodedPassword := formatDatabaseConnectionError(encodedPasswordURL, fmt.Errorf("failed with p@ss:word and p%%40ss%%3Aword")).Error()
	requireContains(encodedPassword, "[REDACTED]")
	requireNotContains(encodedPassword, "p@ss:word")
	requireNotContains(encodedPassword, "p%40ss%3Aword")
}
`

	programPath := filepath.Join(t.TempDir(), "main.go")
	if err := os.WriteFile(programPath, []byte(program), 0o644); err != nil {
		t.Fatal(err)
	}

	command := exec.Command("go", "run", programPath)
	command.Dir = ".."
	output, err := command.CombinedOutput()
	if err != nil {
		t.Fatalf("generated database error behavior failed: %v\n%s", err, output)
	}
}
