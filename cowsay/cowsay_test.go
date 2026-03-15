package cowsay

import (
	"fmt"
	"log"
	"os/exec"
	"testing"
)

func assertOutput(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		fmt.Println(got)
		fmt.Println(want)
		t.Errorf("got diffewant then want")
	}
}

func TestSay(t *testing.T) {
	t.Run("Single line case", func(t *testing.T) {
		phrase := "oi"
		got := buildBaloon(phrase) + buildCow()
		want := executeCowsay(phrase)

		assertOutput(t, got, want)
	})

	t.Run("Two lines case", func(t *testing.T) {
		phrase := "Lorem ipsum dolor sit amet nullam sodales."
		got := buildBaloon(phrase) + buildCow()
		want := executeCowsay(phrase)

		assertOutput(t, got, want)
	})

	t.Run("Big Word Case", func(t *testing.T) {
		phrase := "Pneumonoultramicroscopicsilicovolcanoconiosis"
		got := buildBaloon(phrase) + buildCow()
		want := executeCowsay(phrase)

		assertOutput(t, got, want)
	})

	t.Run("Multiple lines case", func(t *testing.T) {
		phrase := "Lorem ipsum dolor sit amet, consecteturrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrr adipiscing elit. Quisque eu euismod arcu ligula."
		got := buildBaloon(phrase) + buildCow()
		want := executeCowsay(phrase)

		assertOutput(t, got, want)
	})

	t.Run("Multiple lines case", func(t *testing.T) {
		phrase := "charmosa cheirosa nossa que mulher gostosa. SE TA SOLTEIRA VAMO FICAR DE CASAAAAAAAAAAL"
		got := buildBaloon(phrase) + buildCow()
		want := executeCowsay(phrase)

		assertOutput(t, got, want)
	})
}

func executeCowsay(args string) string {
	cmd := exec.Command("cowsay", args)
	cmd.Dir = "/usr/games"

	bytes, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	return string(bytes)
}
