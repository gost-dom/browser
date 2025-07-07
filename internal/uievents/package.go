//go:generate ../../internal/code-gen/code-gen -g eventTypes -p uievents

// Package uievents create code to generate and dispatch UI events.
//
// Properties
// for the different events, like cancelable and bubbles are handled by code
// generated from web IDL specifications, ensuring "correct" bubble and cancel
// behaviour for events.
package uievents
