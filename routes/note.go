package routes

import (
	"mongodb-starter-api/controllers"
	"mongodb-starter-api/middlewares/validators"

	"github.com/gin-gonic/gin"
)

func NoteRoute(router *gin.RouterGroup, handlers ...gin.HandlerFunc) {
	notes := router.Group("/notes", handlers...)
	{
		notes.POST(
			"",
			validators.CreateNoteValidator(),
			controllers.CreateNewNote,
		)

		notes.GET(
			"",
			validators.GetNotesValidator(),
			controllers.GetNotes,
		)

		notes.GET(
			"/:id",
			validators.PathIdValidator(),
			controllers.GetOneNote,
		)

		notes.PUT(
			"/:id",
			validators.PathIdValidator(),
			validators.UpdateNoteValidator(),
			controllers.UpdateNote,
		)

		notes.DELETE(
			"/:id",
			validators.PathIdValidator(),
			controllers.DeleteNote,
		)
	}
}
