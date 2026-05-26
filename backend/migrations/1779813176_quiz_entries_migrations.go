package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(app core.App) error {
		collection := core.NewBaseCollection("quiz_entries")

		quiz, err := app.FindCollectionByNameOrId("quiz")
		if err != nil {
			return err
		}
		users, err := app.FindCollectionByNameOrId("users")
		if err != nil {
			return err
		}

		collection.Fields.Add(&core.TextField{
			Name:     "text",
			Required: true,
			Max:      2000,
		})

		// relation to the post being commented on
		collection.Fields.Add(&core.RelationField{
			Name:          "quiz",
			Required:      true,
			CollectionId:  quiz.Id,
			MaxSelect:     1,
			CascadeDelete: true, // delete comments when the post is deleted
		})

		// relation to the comment's author
		collection.Fields.Add(&core.RelationField{
			Name:          "author",
			Required:      true,
			CollectionId:  users.Id,
			MaxSelect:     1,
			CascadeDelete: true, // delete comments when the user is deleted
		})

		collection.Fields.Add(&core.JSONField{
			Name:    "answers",
			MaxSize: 0,
		})

		collection.Fields.Add(&core.JSONField{
			Name:    "correct_answers",
			MaxSize: 0,
		})

		collection.Fields.Add(&core.AutodateField{
			Name: "created", OnCreate: true,
		})
		collection.Fields.Add(&core.AutodateField{
			Name: "updated", OnCreate: true, OnUpdate: true,
		})

		collection.ListRule = types.Pointer("") // anyone can read
		collection.ViewRule = types.Pointer("")
		collection.CreateRule = types.Pointer("@request.auth.id != '' && author = @request.auth.id")
		collection.UpdateRule = types.Pointer("author = @request.auth.id")
		collection.DeleteRule = types.Pointer("author = @request.auth.id")

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("quiz_entries")

		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
