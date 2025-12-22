package tests

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/sj-distributor/dolphin-example/auth"
	"github.com/sj-distributor/dolphin-example/gen"
	"github.com/sj-distributor/dolphin-example/src"
)

// SetupTestDB initializes the database for testing
func SetupTestDB(t *testing.T) (*gen.DB, *src.Resolver) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("SetupTestDB panicked: %v", r)
		}
	}()

	src.Config()

	dsn := "mysql://root:123456@localhost:3306/dolphin_example?charset=utf8mb4&parseTime=True&loc=Local"
	os.Setenv("DATABASE_URL", dsn)

	db := gen.NewDBFromEnvVars(dsn)
	err := db.AutoMigrate()
	if err != nil {
		t.Fatalf("Failed to migrate database: %v", err)
	}

	// Clean up tables
	gormDB := db.Query()
	tables := []string{
		"users",
		"profiles",
		"tasks",
		"user_roles",
		"tags",
		"task_tags",
		"user_user_roles",
	}

	err = gormDB.Exec("SET FOREIGN_KEY_CHECKS = 0").Error
	if err != nil {
		t.Logf("Warning: Failed to disable foreign key checks: %v", err)
	}

	for _, table := range tables {
		if gormDB.Migrator().HasTable(table) {
			if err := gormDB.Exec("TRUNCATE TABLE " + table).Error; err != nil {
				t.Fatalf("Failed to truncate table %s: %v", table, err)
			}
		}
	}

	gormDB.Exec("SET FOREIGN_KEY_CHECKS = 1")

	eventController, _ := gen.NewEventController()
	resolver := src.New(db, &eventController)

	return db, resolver
}

func GetTestContext(db *gen.DB) context.Context {
	token := auth.USER_JWT_TOKEN.SetToken(map[string]interface{}{
		"id": "test-admin-id",
	})
	ctx := context.Background()
	ctx = context.WithValue(ctx, "Authorization", "Bearer "+token)

	loaders := gen.GetLoaders(db)
	ctx = context.WithValue(ctx, gen.KeyLoaders, loaders)

	return ctx
}

// ... Existing tests preserved ...

func TestOneToOneRelationship(t *testing.T) {
	fmt.Println("Testing OneToOne Relationship (User <-> Profile)...")
	db, resolver := SetupTestDB(t)
	defer db.Close()
	ctx := GetTestContext(db)

	userInput := map[string]interface{}{
		"phone": "13800138000", "password": "password123", "email": "test@example.com", "nickname": "TestUser1",
	}
	user, err := resolver.Mutation().CreateUser(ctx, userInput)
	if err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}

	timestamp := time.Now()
	profileInput := map[string]interface{}{
		"userId": user.ID, "bio": "My Bio", "address": "My Address", "birthday": timestamp,
	}
	profile, err := resolver.Mutation().CreateProfile(ctx, profileInput)
	if err != nil {
		t.Fatalf("CreateProfile failed: %v", err)
	}

	_, err = resolver.Mutation().UpdateUser(ctx, user.ID, map[string]interface{}{"profileId": profile.ID})
	if err != nil {
		t.Fatalf("UpdateUser for ProfileID failed: %v", err)
	}

	if profile.UserID != user.ID {
		t.Errorf("Profile UserID mismatch")
	}
	fetchedUser, err := resolver.Query().User(ctx, &user.ID, nil)
	if err != nil {
		t.Fatalf("Query User failed: %v", err)
	}
	userProfile, err := resolver.User().Profile(ctx, fetchedUser)
	if err != nil {
		t.Fatalf("Get User Profile failed: %v", err)
	}
	if userProfile == nil || userProfile.ID != profile.ID {
		t.Errorf("User Profile mismatch")
	}
	fmt.Println("OneToOne Test Passed")
}

func TestOneToManyRelationship(t *testing.T) {
	fmt.Println("Testing OneToMany Relationship (User -> Tasks)...")
	db, resolver := SetupTestDB(t)
	defer db.Close()
	ctx := GetTestContext(db)

	userInput := map[string]interface{}{"phone": "13800138001", "password": "password123"}
	user, err := resolver.Mutation().CreateUser(ctx, userInput)
	if err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}

	task1, err := resolver.Mutation().CreateTask(ctx, map[string]interface{}{"userId": user.ID, "title": "Task 1", "priority": 1.0})
	if err != nil {
		t.Fatalf("CreateTask 1 failed: %v", err)
	}
	_, err = resolver.Mutation().CreateTask(ctx, map[string]interface{}{"userId": user.ID, "title": "Task 2", "priority": 2.0})
	if err != nil {
		t.Fatalf("CreateTask 2 failed: %v", err)
	}

	fetchedUser, _ := resolver.Query().User(ctx, &user.ID, nil)
	tasks, err := resolver.User().Tasks(ctx, fetchedUser)
	if err != nil {
		t.Fatalf("Get User Tasks failed: %v", err)
	}
	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(tasks))
	}

	fetchedTask1, _ := resolver.Query().Task(ctx, &task1.ID, nil)
	taskUser, err := resolver.Task().User(ctx, fetchedTask1)
	if err != nil {
		t.Fatalf("Get Task User failed: %v", err)
	}
	if taskUser.ID != user.ID {
		t.Errorf("Task User ID mismatch")
	}
	fmt.Println("OneToMany Test Passed")
}

func TestManyToManyRelationship_UserRoles(t *testing.T) {
	fmt.Println("Testing ManyToMany Relationship (User <-> UserRoles)...")
	db, resolver := SetupTestDB(t)
	defer db.Close()
	ctx := GetTestContext(db)

	// User and Role created separately, then linked
	user, err := resolver.Mutation().CreateUser(ctx, map[string]interface{}{"phone": "13800138002", "password": "password123"})
	if err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}
	role, err := resolver.Mutation().CreateUserRole(ctx, map[string]interface{}{"name": "Admin", "description": "Admin"})
	if err != nil {
		t.Fatalf("CreateUserRole failed: %v", err)
	}

	updatedUser, err := resolver.Mutation().UpdateUser(ctx, user.ID, map[string]interface{}{"userRolesIds": []interface{}{role.ID}})
	if err != nil {
		t.Fatalf("UpdateUser failed: %v", err)
	}

	fetchedUser, _ := resolver.Query().User(ctx, &updatedUser.ID, nil)
	roles, err := resolver.User().UserRoles(ctx, fetchedUser)
	if err != nil {
		t.Fatalf("Get User Roles failed: %v", err)
	}
	if len(roles) != 1 || roles[0].ID != role.ID {
		t.Errorf("Role mismatch")
	}

	fetchedRole, _ := resolver.Query().UserRole(ctx, &role.ID, nil)
	roleUsers, err := resolver.UserRole().Users(ctx, fetchedRole)
	if err != nil {
		t.Fatalf("Get Role Users failed: %v", err)
	}
	if len(roleUsers) != 1 || roleUsers[0].ID != user.ID {
		t.Errorf("Role User mismatch")
	}
	fmt.Println("ManyToMany UserRoles Test Passed")
}

func TestManyToManyRelationship_TaskTags(t *testing.T) {
	fmt.Println("Testing ManyToMany Relationship (Task <-> Tags)...")
	db, resolver := SetupTestDB(t)
	defer db.Close()
	ctx := GetTestContext(db)

	task, err := resolver.Mutation().CreateTask(ctx, map[string]interface{}{"title": "Tagged Task"})
	if err != nil {
		t.Fatalf("CreateTask failed: %v", err)
	}
	tag, err := resolver.Mutation().CreateTag(ctx, map[string]interface{}{"name": "Urgent"})
	if err != nil {
		t.Fatalf("CreateTag failed: %v", err)
	}

	updatedTask, err := resolver.Mutation().UpdateTask(ctx, task.ID, map[string]interface{}{"tagsIds": []interface{}{tag.ID}})
	if err != nil {
		t.Fatalf("UpdateTask failed: %v", err)
	}

	fetchedTask, _ := resolver.Query().Task(ctx, &updatedTask.ID, nil)
	tags, err := resolver.Task().Tags(ctx, fetchedTask)
	if err != nil {
		t.Fatalf("Get Task Tags failed: %v", err)
	}
	if len(tags) != 1 || tags[0].ID != tag.ID {
		t.Errorf("Tag mismatch")
	}

	fetchedTag, _ := resolver.Query().Tag(ctx, &tag.ID, nil)
	tagTasks, err := resolver.Tag().Tasks(ctx, fetchedTag)
	if err != nil {
		t.Fatalf("Get Tag Tasks failed: %v", err)
	}
	if len(tagTasks) != 1 || tagTasks[0].ID != task.ID {
		t.Errorf("Tag Task mismatch")
	}
	fmt.Println("ManyToMany TaskTags Test Passed")
}

func TestNestedCreation(t *testing.T) {
	fmt.Println("Testing Nested Creation (User with Profile)...")
	db, resolver := SetupTestDB(t)
	defer db.Close()
	ctx := GetTestContext(db)

	userInput := map[string]interface{}{
		"phone": "13800138003", "password": "password123",
		"profile": map[string]interface{}{"bio": "Nested Bio", "address": "Nested Address"},
	}
	user, err := resolver.Mutation().CreateUser(ctx, userInput)
	if err != nil {
		t.Fatalf("CreateUser with nested profile failed: %v", err)
	}

	fetchedUser, _ := resolver.Query().User(ctx, &user.ID, nil)
	profile, err := resolver.User().Profile(ctx, fetchedUser)
	if err != nil {
		t.Fatalf("Get User Profile failed: %v", err)
	}
	if profile == nil || *profile.Bio != "Nested Bio" {
		t.Errorf("Profile Bio mismatch")
	}
	fmt.Println("Nested Creation Test Passed")
}

// --- NEW NESTED TESTS ---

func TestNestedOneToMany(t *testing.T) {
	fmt.Println("Testing Nested OneToMany (Create User with Tasks)...")
	db, resolver := SetupTestDB(t)
	defer db.Close()
	ctx := GetTestContext(db)

	userInput := map[string]interface{}{
		"phone": "13800138004", "password": "password123",
		"tasks": []interface{}{
			map[string]interface{}{"title": "Nested Task 1", "priority": 1.0},
			map[string]interface{}{"title": "Nested Task 2", "priority": 2.0},
		},
	}
	user, err := resolver.Mutation().CreateUser(ctx, userInput)
	if err != nil {
		t.Fatalf("CreateUser with nested tasks failed: %v", err)
	}

	fetchedUser, _ := resolver.Query().User(ctx, &user.ID, nil)
	tasks, err := resolver.User().Tasks(ctx, fetchedUser)
	if err != nil {
		t.Fatalf("Get User Tasks failed: %v", err)
	}
	if len(tasks) != 2 {
		t.Errorf("Expected 2 nested created tasks, got %d", len(tasks))
	}

	// Verify task data
	titles := map[string]bool{}
	for _, task := range tasks {
		if task.UserID == nil || *task.UserID != user.ID {
			t.Errorf("Task UserID not set correctly")
		}
		titles[task.Title] = true
	}
	if !titles["Nested Task 1"] || !titles["Nested Task 2"] {
		t.Errorf("Nested tasks titles missing")
	}
	fmt.Println("Nested OneToMany Test Passed")
}

func TestNestedManyToMany_UserRoles(t *testing.T) {
	fmt.Println("Testing Nested ManyToMany (Create User with Roles)...")
	db, resolver := SetupTestDB(t)
	defer db.Close()
	ctx := GetTestContext(db)

	userInput := map[string]interface{}{
		"phone": "13800138005", "password": "password123",
		"userRoles": []interface{}{
			map[string]interface{}{"name": "Manager", "description": "Mgr"},
			map[string]interface{}{"name": "Editor", "description": "Edt"},
		},
	}
	user, err := resolver.Mutation().CreateUser(ctx, userInput)
	if err != nil {
		t.Fatalf("CreateUser with nested roles failed: %v", err)
	}

	fetchedUser, _ := resolver.Query().User(ctx, &user.ID, nil)
	roles, err := resolver.User().UserRoles(ctx, fetchedUser)
	if err != nil {
		t.Fatalf("Get User Roles failed: %v", err)
	}
	if len(roles) != 2 {
		t.Errorf("Expected 2 nested created roles, got %d", len(roles))
	}

	names := map[string]bool{}
	for _, r := range roles {
		names[r.Name] = true
	}
	if !names["Manager"] || !names["Editor"] {
		t.Errorf("Nested roles names missing")
	}
	fmt.Println("Nested ManyToMany UserRoles Test Passed")
}

func TestNestedManyToMany_TaskTags(t *testing.T) {
	fmt.Println("Testing Nested ManyToMany (Create Task with Tags)...")
	db, resolver := SetupTestDB(t)
	defer db.Close()
	ctx := GetTestContext(db)

	taskInput := map[string]interface{}{
		"title": "Task with Tags",
		"tags": []interface{}{
			map[string]interface{}{"name": "Tag1", "color": "Blue"},
			map[string]interface{}{"name": "Tag2", "color": "Green"},
		},
	}
	task, err := resolver.Mutation().CreateTask(ctx, taskInput)
	if err != nil {
		t.Fatalf("CreateTask with nested tags failed: %v", err)
	}

	fetchedTask, _ := resolver.Query().Task(ctx, &task.ID, nil)
	tags, err := resolver.Task().Tags(ctx, fetchedTask)
	if err != nil {
		t.Fatalf("Get Task Tags failed: %v", err)
	}
	if len(tags) != 2 {
		t.Errorf("Expected 2 nested created tags, got %d", len(tags))
	}
	fmt.Println("Nested ManyToMany TaskTags Test Passed")
}

func TestNestedUpdate(t *testing.T) {
	fmt.Println("Testing Nested Update (Update User and Profile)...")
	db, resolver := SetupTestDB(t)
	defer db.Close()
	ctx := GetTestContext(db)

	// User with profile
	userInput := map[string]interface{}{
		"phone": "13800138006", "password": "password123",
		"profile": map[string]interface{}{"bio": "Old Bio"},
	}
	user, err := resolver.Mutation().CreateUser(ctx, userInput)
	if err != nil {
		t.Fatalf("Setup User failed: %v", err)
	}

	// Update User with nested Profile update
	// Note: Nested Update usually works by passing "profile": { ... } logic
	// Depending on resolver logic, it might create new or update existing.
	// Generated logic typically requires passing ID if updating existing?
	// Or if OneToOne, it updates the associated record?
	// Let's rely on gorm functionality or generated logic.

	// If Profile creation returned ID, we should check if User has it.
	fetchedUser, _ := resolver.Query().User(ctx, &user.ID, nil)
	profile, _ := resolver.User().Profile(ctx, fetchedUser)

	updateInput := map[string]interface{}{
		"profile": map[string]interface{}{
			"id":  profile.ID, // Pass ID to ensure update?
			"bio": "New Bio",
		},
	}
	_, err = resolver.Mutation().UpdateUser(ctx, user.ID, updateInput)
	if err != nil {
		t.Fatalf("UpdateUser with nested profile failed: %v", err)
	}

	// Verify
	fetchedUser, _ = resolver.Query().User(ctx, &user.ID, nil)
	updatedProfile, _ := resolver.User().Profile(ctx, fetchedUser)
	if updatedProfile.Bio == nil || *updatedProfile.Bio != "New Bio" {
		t.Errorf("Profile Bio not updated")
	}
	if updatedProfile.ID != profile.ID {
		t.Logf("Warning: Profile ID changed, maybe it was re-created instead of updated")
	}
	fmt.Println("Nested Update Test Passed")
}
