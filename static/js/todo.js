var app = angular.module('myApp', []);
app.controller('myTask', function($scope,$http) {
    $scope.tasks = [];
     var refresh = function() {
    return $http.get('/task/').
      success(function(data) { $scope.tasks = data.Tasks; })
  };



   $scope.todo = function(){
     $http.post('/task', {Title: $scope.todolist})
     .success(function() {
        refresh().then(function() {
     
          $scope.todoText = '';
        })
      });
  

    }

});