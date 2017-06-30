<!DOCTYPE html>
<html>
<head>
  <title>ToDoList</title>
  <script type="text/javascript" src="/static/js/angular.min.js"></script>
  <script type="text/javascript" src="/static/js/todo.js"></script>

</head>
<body>
    <div ng-app="myApp" ng-controller="myTask">
      <h1>ToDo List</h1>
      <ul class='grey rounded-box'>
        <li ng-repeat='t in tasks'>
        <span class='checkbox'></span>{{t.Title}}
        </li>
      </ul>
        <form>
          <input type="text" name="task" ng-model="todolist">
          <button class="submit" ng-click="todo()">上传</button>
       </form>
    </div>

</body>
</html>