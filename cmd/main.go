package main

import (
	"log/slog"

	"user-service/internal/repository"
	"user-service/internal/service"
)

func main() {

	slog.Info("--- Using InMemoryUserRepo ---")
	demonstrateUserService(repository.NewInMemoryUserRepo())

	slog.Info("--- Using MockUserRepo ---")
	demonstrateUserService(repository.NewMockUserRepo())
}

func demonstrateUserService(userRepo repository.UserRepository) {
	userService := service.NewUserService(userRepo)

	user1, err := userService.CreateUser("John Doe", "john@example.com", "admin")
	if err != nil {
		slog.Error("Error creating user1", "error", err)
		return
	}
	slog.Info("Created user", "user", user1)

	user2, err := userService.CreateUser("Jane Smith", "jane@example.com", "user")
	if err != nil {
		slog.Error("Error creating user2", "error", err)
		return
	}
	slog.Info("Created user", "user", user2)

	user3, err := userService.CreateUser("Bob Guest", "bob@example.com", "guest")
	if err != nil {
		slog.Error("Error creating user3", "error", err)
		return
	}
	slog.Info("Created user", "user", user3)

	_, err = userService.CreateUser("Invalid User", "invalid@example.com", "invalid_role")
	if err != nil {
		slog.Warn("Expected error for invalid role", "error", err)
	}

	slog.Info("All users:")
	users := userService.ListUsers()
	for i, user := range users {
		slog.Info("User", "index", i+1, "user", user)
	}

	slog.Info("Getting user by ID", "userID", user1.ID)
	retrievedUser, err := userService.GetUser(user1.ID)
	if err != nil {
		slog.Error("Error getting user", "error", err)
	} else {
		slog.Info("Retrieved user", "user", retrievedUser)
	}

	_, err = userService.GetUser("non-existent-id")
	if err != nil {
		slog.Warn("Expected error for non-existent user", "error", err)
	}

	slog.Info("Removing user", "userID", user2.ID)
	err = userService.RemoveUser(user2.ID)
	if err != nil {
		slog.Error("Error removing user", "error", err)
	} else {
		slog.Info("User removed successfully")
	}

	slog.Info("Users after removal:")
	users = userService.ListUsers()
	for i, user := range users {
		slog.Info("User", "index", i+1, "user", user)
	}
}
