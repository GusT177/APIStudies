package main

import(
  "net/http"
  "github.com/gin-gonic/gin"
)

type infected struct {
  
  ID string `json:"id"`
  Name string `json:"name"`
  Skill string `json:"skill"`
  Skill2 string `json:"skill2"`  


}

var infecteds = []infected{
  {ID: "1", Name: "hunter", Skill: "Jump High", Skill2: "Scratch" },
  {ID: "2", Name: "boomer", Skill: "Pilk", Skill2: "Scratch"},
  {ID: "3", Name: "smoker", Skill: "Tongue Squeeze", Skill2: "Scratch"},
  {ID: "4", Name: "tank", Skill: "Throw Rocks", Skill2: "Super punch"},

}

func GetInfecteds(c *gin.Context){
  c.IndentedJSON(http.StatusOK, infecteds)

}

func PostInfecteds(c *gin.Context){
  var NewInfected infected

  if err := c.BindJSON(&NewInfected); err != nil {return}

  infecteds = append(infecteds, NewInfected)
  c.IndentedJSON(http.StatusCreated, NewInfected)

}


func main(){
  router := gin.Default()
  router.GET("/infected", GetInfecteds)
  router.POST("/infected", PostInfecteds)
  router.Run("localhost:8080")
}
