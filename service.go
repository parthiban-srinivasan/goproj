package main

import (
         "errors"
         "strings"
)

type StringService interface {
         Uppercase(string) (string, error)
         Count(string) int
}

type stringService struct{}

func (stringService) Uppercase(inStr string) (string, error) {
         if inStr == "" {
                 return "", ErrEmpty         
         }
         return strings.ToUpper(inStr), nil
}

func (stringService) Count(inStr string) int {
         return len(inStr)
}

var ErrEmpty = errors.New("Empty String")
