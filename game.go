package main

import (
	"fmt"

	mid "github.com/middangeard-fiction/middangeard"
)

var game mid.Game

// Items
var cloak mid.Item
var hook mid.Item
var message mid.Item

func newItem(test string) mid.Item {
	test2 := new(mid.Item)
	return *test2
}

func init() {
	game = mid.Game{
		Title:  "Cloak of Darkness",
		Author: "Roger Firth (implemented by Mark Bauermeister)",
		Intro: `Hurrying through the rainswept November night, you're glad to see the
		bright lights of the Opera House.
		
		It's surprising that there aren't more
		people about but, hey, what do you expect in a cheap demo game...?`,
		Goodbye:  "See you later!",
		MaxScore: 2,
		Verbs: map[string]func(...string){
			"drop": drop,
		},
		Synonyms: map[string][]string{
			"drop": {"throw"},
		},
	}

	mid.Rankings = []mid.Ranking{
		{Score: 0, Rank: "Bloody Beginner"},
		{Score: game.MaxScore / 2, Rank: "Amateur"},
		{Score: game.MaxScore, Rank: "Master Adventurer"},
	}

	game.Items = map[*mid.Item]*mid.Item{
		&cloak: {
			Article:  "a", // What if we want to use "the" in other parts though? Hmmm...
			Name:     "velvet cloak",
			Synonyms: []string{"cloak", "dark", "black", "satin", "velvet"},
			Description: `A handsome cloak, of velvet trimmed with satin, and slightly
			spattered with raindrops. Its blackness is so deep that it
			almost seems to suck light from the room.`,
			Carryable: true,
			Verbs: mid.ItemVerbs{
				"drop": func(item *mid.Item, room *mid.Room) {
					if game.Player.Location == "cloakroom" {
						game.Player.AwardPoints(1)
						game.Rooms["bar"].Lit = true
					} else {
						game.Output(`This isn't the best place to leave a smart cloak
						lying around.`)
						game.Player.PickupItem(game.Items[&cloak])
						room.Items.Remove(game.Items[&cloak])
					}
				},
				"take": func(item *mid.Item, room *mid.Room) {
					game.Rooms["bar"].Lit = false
				},
			},
		},
		&hook: {
			Article:     "a",
			Name:        "small brass hook",
			Synonyms:    []string{"small", "brass", "hook", "peg"},
			Description: `It's just a small brass hook, screwed to the wall.`,
			Fixture:     true,
		},
		&message: {
			Article:  "a",
			Name:     "scrawled message",
			Synonyms: []string{"message", "floor", "sawdust"},
			Verbs: mid.ItemVerbs{
				"inspect": func(item *mid.Item, room *mid.Room) {
					if !room.Lit {
						game.Output(`In the dark? You could easily disturb something!`)
					} else {
						game.Player.AwardPoints(1)
						game.Output(`The message, neatly marked in the sawdust, reads...`)
					}
				},
			},
		},
	}

	game.Player = mid.Player{
		Name:        "Self",
		Description: "Just an average individual.",
		Location:    "foyer",
		Score:       0,
		Inventory: mid.Items{
			game.Items[&cloak],
		},
	}

	game.Rooms = map[string]*mid.Room{
		"foyer": {
			Name: "Foyer of the Opera House",
			Description: `You are standing in a spacious hall, splendidly decorated in red
            and gold, with glittering chandeliers overhead. The entrance from
			the street is to the north, and there are doorways south and west.`,
			Lit: true,
			OnEnter: func(r *mid.Room) {
				if r.Visited {
					r.Description = "Nothing has changed since you last entered here."
				}
			},
			Directions: mid.Directions{
				West:  "cloakroom",
				South: "bar",
				North: `You've only just arrived, and besides, the weather outside
				seems to be getting worse.`,
			},
		},
		"cloakroom": {
			Name: "Cloakroom",
			Description: `The walls of this small room were clearly once lined with hooks, though now only one remains.
			The exit is a door to the east.`,
			Lit: true,
			Directions: mid.Directions{
				East: "foyer",
			},
			Items: mid.Items{
				game.Items[&hook],
			},
		},
		"bar": {
			Name: "Foyer Bar",
			Description: `The bar, much rougher than you'd have guessed
			after the opulence of the foyer to the north, is completely empty.
			There seems to be some sort of message scrawled in the sawdust on the floor.`,
			Directions: mid.Directions{
				North: "foyer",
			},
			OnEnter: func(r *mid.Room) {
				if r.Lit {
					fmt.Println("Room is lit")
				}
			},
			Items: mid.Items{
				game.Items[&message],
			},
		},
	}
}

func drop(args ...string) {
	fmt.Println("You dropped something")
}
