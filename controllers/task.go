package controllers

import (
    "github.com/astaxie/beego"
    "quickstart/models"
    "encoding/json"
)

type TaskController struct {
    beego.Controller
}


func (c *TaskController) GetTask() {
    
    res := struct{ Tasks []*models.Task }{models.DefaultTaskList.All()}
    c.Data["json"] = res
    c.ServeJSON()
}


func (c *TaskController) AddTask() {
   ob :=struct { Title string} { }

    if err := json.Unmarshal(c.Ctx.Input.RequestBody, &ob); err != nil {
        c.Ctx.Output.SetStatus(400)
        c.Ctx.Output.Body([]byte("empty title"))
        return
    }
    t, err := models.NewTask(ob.Title)
    if err != nil {
        c.Ctx.Output.SetStatus(400)
        c.Ctx.Output.Body([]byte(err.Error()))
        return
    }
     models.DefaultTaskList.Save(t)
    

}
