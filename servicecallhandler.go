package servicecallhandler

import (  "fmt"
  "encoding/json"
  "io/ioutil"
  "net/http"
  "os"
  "log"
  "strconv"
  "bufio"
  "strings")


type Cocktails struct {
 Cocktails []Cocktail `json:"drinks"`

}

type Cocktail struct {
 Id int `json:"idDrink"`
 Name string `json:"strDrink"`
 Alchohlic string `json:"strAlcoholic"`
 IngridentOne string `json:"strIngredient1"`
 IngridentTwo string `json:"strIngredient2"`
 IngridentThree string `json:"strIngredient3"`
 IngridentFour string `json:"strIngredient4"`
 IngridentFive string `json:"strIngredient5"`
 Instructions string `json:"strInstructions"`
 Thumbnail string `json:"strDrinkThumb"`

}

func ingridents(c Cocktail ) string {
 var v = "\n" +c.IngridentOne+ "\n" +c.IngridentTwo+ "\n"+c.IngridentThree+ "\n " + c.IngridentFour + "\n" + c.IngridentFive
  return v
}



func main() {


    
}
func cocktailCall (){
  //will replace with the properties call instead when given the button input
response, err := http.Get("https://www.thecocktaildb.com/api/json/v1/1/search.php?s=margarita");

if err != nil {
   fmt.Print(err.Error())
        os.Exit(1)
}


data, err := ioutil.ReadAll(response.Body)
if err != nil {
     fmt.Print(err.Error())
  log.Fatal(err)
}
//fmt.Println(string(data))
  var cock Cocktails

json.Unmarshal(data, &cock)

for i:= 0; i < len(cock.Cocktails);i++ {
fmt.Println("Name: ",cock.Cocktails[i].Name,"\n",
"Non/Alcholic: ",cock.Cocktails[i].Alchohlic ,"\n","Instructions: ",cock.Cocktails[i].Instructions,"\n","Ingridents: ",ingridents(cock.Cocktails[i]));
}

}
