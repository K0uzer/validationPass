package main

// ------------
// ------------ Тесты (отдельный файл _test.go) ------------
// ------------

import "testing"
 
func TestPasswordValidator(t *testing.T) {
    testCases := []struct {
        name      string
        password  string
        validator validator
        expected  bool
    }{
        {
            name:      "validator checks if password contains at least one digit, but it doesn't",
            password:  "password",
            validator: digits,
            expected:  false,
        },
        {
            name:      "validator checks if password contains at least one digit, and it does",
            password:  "password123",
            validator: digits,
            expected:  true,
        },
        {
            name:      "unicode digit",
            password:  "пароль٣",
            validator: digits,
            expected:  true,
        },
        {
            name:      "validator checks if password contains at least one letter, but it doesn't",
            password:  "123456",
            validator: letters,
            expected:  false,
        },
        {
            name:      "validator checks if password contains at least one letter, and it does",
            password:  "password123",
            validator: letters,
            expected:  true,
        },
        {
            name:      "unicode letter",
            password:  "123Ж",
            validator: letters,
            expected:  true,
        },
        {
            name:      "minlen counts runes, not bytes",
            password:  "абв",
            validator: minlen(3),
            expected:  true,
        },
        {
            name:      "check if password contains at least one digit and one letter, and it does",
            password:  "password123",
            validator: and(digits, letters),
            expected:  true,
        },
        {
            name:      "check if password contains at least one digit OR one letter, and it contains only digits",
            password:  "123456",
            validator: or(digits, letters),
            expected:  true,
        },
        {
            name:      "check if password contains at least one digit and one letter and has length at least 8",
            password:  "password123",
            validator: and(minlen(8), digits, letters),
            expected:  true,
        },
        {
            name:      "check if password contains at least one digit and one letter and has length at least 8, but it doesn't have a digit",
            password:  "password",
            validator: and(minlen(8), digits, letters),
            expected:  false,
        },
        {
            name:      "check if password contains at least one digit and one letter and has length at least 8, but it doesn't have a letter",
            password:  "12345678",
            validator: and(minlen(8), digits, letters),
            expected:  false,
        },
        {
            name:      "and without validators is true",
            password:  "anything",
            validator: and(),
            expected:  true,
        },
        {
            name:      "or without validators is false",
            password:  "anything",
            validator: or(),
            expected:  false,
        },
    }
 
    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            actual := tc.validator(tc.password)
            if actual != tc.expected {
                t.Errorf("Expected %v, got %v", tc.expected, actual)
            }
        })
    }
}