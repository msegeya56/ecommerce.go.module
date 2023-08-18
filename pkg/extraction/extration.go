 package modelstoentities

// import "github.com/msegeya56/ecommerce.go.module/pkg/domains/models"

// type ExtractEntitiesToModels struct {
// 	Cloner *entitiestomodels.CloneEntitiesToModels
// }

// func New() *ExtractEntitiesToModels {

// 	return &ExtractEntitiesToModels{
// 		Cloner: entitiestomodels.New(),
// 	}
// }


// // Clone Models from Entity

// func (e *ExtractEntitiesToModels) CloneAccessModelFromEntity(ientity interface{}) models.Customer {

// 	cloned, err := e.Cloner.EntityToModels(ientity)

// 	if err != nil {

// 		panic(err)
// 	}

// 	return cloned.(models.Access)
// }

// func (e *ExtractEntitiesToModels) ClonePointerAccessModelFromEntity(ientity interface{}) *models.Customer {

// 	cloned, err := e.Cloner.EntityToModels(ientity)

// 	if err != nil {

// 		panic(err)
// 	}

// 	return cloned.(*models.Customer)
// }