package note

import (
	"context"
	"fmt"
	"log"

	"github.com/SojebSikder/goframe/ent"
)

// create record
func Create(ctx context.Context, client *ent.Client, title string, content string) *ent.Note {
	u, err := client.Note.
		Create().
		SetTitle(title).
		SetContent(content).
		Save(ctx)
	if err != nil {
		return nil
	}
	log.Println("note was created: ", u)
	return nil
}

func FindAll(ctx context.Context, client *ent.Client) (any, error) {
	u, err := client.Note.
		Query().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying note: %w", err)
	}
	log.Println("note returned: ", u)
	return u, nil
}

func FindOne(ctx context.Context, client *ent.Client) (any, error) {
	u, err := client.Note.
		Query().
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying note: %w", err)
	}
	log.Println("note returned: ", u)
	return u, nil
}

func Update(ctx context.Context, client *ent.Client, title string, content string) *ent.Note {
	u, err := client.Note.
		Update().
		SetTitle(title).
		SetContent(content).
		Save(ctx)
	if err != nil {
		return nil
	}
	log.Println("note was created: ", u)
	return nil
}

func Delete(ctx context.Context, client *ent.Client, id int) error {
	err := client.Note.
		DeleteOneID(id).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed querying note: %w", err)
	}
	return nil
}
