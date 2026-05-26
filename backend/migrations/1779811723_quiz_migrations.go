package migrations

import (
	"github.com/pocketbase/pocketbase/tools/types"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {

		users, err := app.FindCollectionByNameOrId("users")
		if err != nil {
			return err
		}

		collection := core.NewBaseCollection("quiz")

		collection.ListRule = types.Pointer("")
		collection.ViewRule = types.Pointer("")
		collection.CreateRule = types.Pointer("@request.auth.id != '' && author = @request.auth.id") // authed users
		collection.UpdateRule = types.Pointer("@request.auth.id != '' && author = @request.auth.id")
		collection.DeleteRule = nil

		collection.Fields.Add(&core.TextField{
			Name:     "name",
			Required: true,
			Max:      200,
		})

		collection.Fields.Add(&core.TextField{
			Name:     "description",
			Required: true,
			Max:      200,
		})

		collection.Fields.Add(&core.RelationField{
			Name:          "author",
			Required:      true, // every post must have an owner
			CollectionId:  users.Id,
			MaxSelect:     1,    // single user, not many
			CascadeDelete: true, // delete posts when the user is deleted
		})

		collection.Fields.Add(&core.SelectField{
			Name:      "visibility",
			Values:    []string{"public_auth", "public_noauth", "group", "unlisted"},
			Required:  true,
			MaxSelect: 1,
		})

		collection.Fields.Add(&core.BoolField{
			Name:     "is_deleted",
			Required: true,
		})

		collection.Fields.Add(&core.AutodateField{
			Name:     "created",
			OnCreate: true,
		})
		collection.Fields.Add(&core.AutodateField{
			Name:     "updated",
			OnCreate: true,
			OnUpdate: true,
		})

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("quiz")

		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
