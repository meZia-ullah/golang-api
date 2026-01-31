package db

import (
	"time"

	"goapi/models"
)

func SeedUsers() {
	users := []models.User{
		{
			Username: "john_doe",
			Email:    "john@example.com",
			Password: "password123",
			IsActive: true,
		},
		{
			Username: "jane_smith",
			Email:    "jane@example.com",
			Password: "password123",
			IsActive: true,
		},
		{
			Username: "bob_wilson",
			Email:    "bob@example.com",
			Password: "password123",
			IsActive: true,
		},
		{
			Username: "alice_brown",
			Email:    "alice@example.com",
			Password: "password123",
			IsActive: true,
		},
		{
			Username: "charlie_davis",
			Email:    "charlie@example.com",
			Password: "password123",
			IsActive: false,
		},
	}

	for _, user := range users {
		DB.FirstOrCreate(&user, models.User{Email: user.Email})
	}
}

func SeedStaff() {
	staffMembers := []models.Staff{
		{
			ShopID:    1,
			RoleID:    1,
			FirstName: "Ahmad",
			LastName:  "Jawabreh",
			Email:     "ahmad@gmail.com",
			Phone:     strPtr("0791234567"),
			Password:  "hashed",
			IsActive:  true,
		},
		{
			ShopID:    1,
			RoleID:    2,
			FirstName: "Fatima",
			LastName:  "Hassan",
			Email:     "fatima@gmail.com",
			Phone:     strPtr("0792345678"),
			Password:  "hashed",
			IsActive:  true,
		},
		{
			ShopID:    1,
			RoleID:    3,
			FirstName: "Mohammed",
			LastName:  "Ali",
			Email:     "mohammed@gmail.com",
			Phone:     strPtr("0793456789"),
			Password:  "hashed",
			IsActive:  true,
		},
		{
			ShopID:    2,
			RoleID:    1,
			FirstName: "Sarah",
			LastName:  "Johnson",
			Email:     "sarah@gmail.com",
			Phone:     strPtr("0794567890"),
			Password:  "hashed",
			IsActive:  true,
		},
		{
			ShopID:    2,
			RoleID:    2,
			FirstName: "David",
			LastName:  "Smith",
			Email:     "david@gmail.com",
			Phone:     strPtr("0795678901"),
			Password:  "hashed",
			IsActive:  true,
		},
		{
			ShopID:    2,
			RoleID:    3,
			FirstName: "Emily",
			LastName:  "Brown",
			Email:     "emily@gmail.com",
			Phone:     strPtr("0796789012"),
			Password:  "hashed",
			IsActive:  false,
		},
	}

	for _, staff := range staffMembers {
		DB.FirstOrCreate(&staff, models.Staff{Email: staff.Email})
	}
}

func SeedTodos() {
	dueDate1 := time.Now().AddDate(0, 0, 7)
	dueDate2 := time.Now().AddDate(0, 0, 14)
	dueDate3 := time.Now().AddDate(0, 1, 0)

	todos := []models.Todo{
		{
			UserID:      1,
			Title:       "Complete project proposal",
			Description: "Finish the Q1 project proposal and send to manager",
			IsCompleted: false,
			DueDate:     &dueDate1,
		},
		{
			UserID:      1,
			Title:       "Review code",
			Description: "Review pull requests from team members",
			IsCompleted: true,
			DueDate:     &dueDate2,
		},
		{
			UserID:      1,
			Title:       "Update documentation",
			Description: "Update API documentation with new endpoints",
			IsCompleted: false,
			DueDate:     &dueDate3,
		},
		{
			UserID:      2,
			Title:       "Buy groceries",
			Description: "Milk, eggs, bread, and vegetables",
			IsCompleted: false,
			DueDate:     &dueDate1,
		},
		{
			UserID:      2,
			Title:       "Call dentist",
			Description: "Schedule appointment for dental checkup",
			IsCompleted: true,
			DueDate:     &dueDate2,
		},
		{
			UserID:      2,
			Title:       "Finish reading book",
			Description: "Complete the novel by next week",
			IsCompleted: false,
			DueDate:     &dueDate3,
		},
		{
			UserID:      3,
			Title:       "Prepare presentation",
			Description: "Create slides for team meeting on Friday",
			IsCompleted: false,
			DueDate:     &dueDate1,
		},
		{
			UserID:      3,
			Title:       "Fix bug in login page",
			Description: "Debug the authentication issue reported by users",
			IsCompleted: true,
			DueDate:     &dueDate2,
		},
		{
			UserID:      3,
			Title:       "Schedule team lunch",
			Description: "Reserve restaurant for team building event",
			IsCompleted: false,
			DueDate:     &dueDate3,
		},
		{
			UserID:      4,
			Title:       "Gym session",
			Description: "30 min cardio and weight training",
			IsCompleted: true,
			DueDate:     &dueDate1,
		},
		{
			UserID:      4,
			Title:       "Pay bills",
			Description: "Electricity, water, and internet bills",
			IsCompleted: false,
			DueDate:     &dueDate2,
		},
		{
			UserID:      4,
			Title:       "Car maintenance",
			Description: "Oil change and tire rotation",
			IsCompleted: false,
			DueDate:     &dueDate3,
		},
		{
			UserID:      1,
			Title:       "Client meeting",
			Description: "Discuss project requirements with client",
			IsCompleted: false,
			DueDate:     &dueDate1,
		},
		{
			UserID:      2,
			Title:       "Laundry day",
			Description: "Wash and fold clothes",
			IsCompleted: true,
			DueDate:     &dueDate2,
		},
		{
			UserID:      3,
			Title:       "Deploy to production",
			Description: "Release v2.0 to production server",
			IsCompleted: false,
			DueDate:     &dueDate3,
		},
	}

	for _, todo := range todos {
		DB.FirstOrCreate(&todo, models.Todo{UserID: todo.UserID, Title: todo.Title})
	}
}

func strPtr(s string) *string {
	return &s
}

func Seed() {
	SeedUsers()
	SeedStaff()
	SeedTodos()
}
