package dm

import (
	"gorm.io/gorm"
	"inventory/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{
		db: db,
	}
}

func (h *Handler) ViewAllHandler(g *gin.Context) {
	data := &service.Inventory{}
	inv := service.NewInventoryOrm(h.db, data)

	invList, err := inv.GetAll()
	if err != nil {
		log.Printf("err:%v", err)
		g.JSON(http.StatusInternalServerError, gin.H{"data": nil, "debugMsg": "Error"})
		return
	}
	g.JSON(http.StatusOK, gin.H{"data": invList, "debugMsg": ""})
	return
}

func (h *Handler) CreateHandle(g *gin.Context) {
	data := &service.Inventory{}
	err := g.ShouldBindJSON(data)
	if err != nil {
		log.Printf("err:%v", err)
		g.JSON(http.StatusUnprocessableEntity, gin.H{"data": nil, "debugMsg": "Invalid json provided"})
		return
	}
	log.Printf("data:%v", data)
	inv := service.NewInventoryOrm(h.db, data)
	err = inv.Create()
	if err != nil {
		log.Printf("err:%v", err)
		g.JSON(http.StatusInternalServerError, gin.H{"data": nil, "debugMsg": "Error"})
		return
	}
	g.String(http.StatusOK, "Appliance added successfully", nil)
	return
}

func (h *Handler) SearchHandle(g *gin.Context) {
	data := &service.Inventory{}
	err := g.ShouldBindJSON(data)
	if err != nil {
		log.Printf("err:%v", err)
		g.JSON(http.StatusUnprocessableEntity, gin.H{"data": nil, "debugMsg": "Invalid json provided"})
		return
	}
	inv := service.NewInventoryOrm(h.db, data)
	log.Printf("data:%v", data)
	invList, err := inv.GetByWhere()
	if err != nil {
		log.Printf("err:%v", err)
		g.JSON(http.StatusInternalServerError, gin.H{"data": nil, "debugMsg": "Error"})
		return
	}
	g.JSON(http.StatusOK, gin.H{"data": invList, "debugMsg": ""})
	return
}

func (h *Handler) UpdateHandle(g *gin.Context) {
	data := &service.Inventory{}
	err := g.ShouldBindJSON(data)
	if err != nil {
		log.Printf("err:%v", err)
		g.JSON(http.StatusUnprocessableEntity, gin.H{"data": nil, "debugMsg": "Invalid json provided"})
		return
	}
	inv := service.NewInventoryOrm(h.db, data)
	log.Printf("data:%v", data)
	err = inv.Update(data.Id)
	if err != nil {
		log.Printf("err:%v", err)
		g.JSON(http.StatusInternalServerError, gin.H{"data": nil, "debugMsg": "Error"})
		return
	}
	g.JSON(http.StatusOK, gin.H{"data": inv, "debugMsg": ""})
	return
}

func (h *Handler) DeleteHandle(g *gin.Context) {
	data := &service.Inventory{}
	err := g.ShouldBindJSON(data)
	if err != nil {
		log.Printf("err:%v", err)
		g.JSON(http.StatusUnprocessableEntity, gin.H{"data": nil, "debugMsg": "Invalid json provided"})
		return
	}
	inv := service.NewInventoryOrm(h.db, data)
	log.Printf("data:%v", data)
	err = inv.Delete()
	if err != nil {
		log.Printf("err:%v", err)
		g.JSON(http.StatusInternalServerError, gin.H{"data": nil, "debugMsg": "Error"})
		return
	}
	g.JSON(http.StatusOK, gin.H{"data": inv, "debugMsg": ""})
	return
}

//type Processor struct {
//	service service.InventoryInt
//}
//
//func (p *Processor) ViewAll() (err error) {
//
//	g.JSON(http.StatusOK, gin.H{
//		"data":     u,
//		"debugMsg": "",
//	})
//}
