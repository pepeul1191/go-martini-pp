function RootCtrl ($scope, $http, $location, $window, $route) {

    if (sessionStorage.is_user_logged_in == undefined)
        $scope.is_user_logged_in = false;
    else
        $scope.is_user_logged_in = sessionStorage.is_user_logged_in;

    $scope.show_login_error  = false;

    $scope.username = sessionStorage.username;

    //
    // Login and Logout handlers
    //
    $scope.login = function(username, password) {
        if (username == undefined || password == undefined){
            $scope.login_error      = "Username and password can't be empty";
            $scope.show_login_error = true;
            return;
        }

        // clear session storage for new session
        sessionStorage.clear();

        // save new username
        sessionStorage.username = username;

        $scope.username = username;

        //
        // Execute login
        //
        $http.post("/login", {"username" : username, "password" : password}).success(function(response) {
                if (response.error == undefined){
                    $location.path('/dashboard');
                    $scope.is_user_logged_in = true;
                    sessionStorage.is_user_logged_in = true;
                }
                else {
                    $scope.login_error = response.error.message;
                    $scope.show_login_error = true;
                }
            }
        );
    }
    
    $scope.logout = function() {
        $http.post("/logout", {"username" : sessionStorage.username}).success(function(response){
            $scope.is_user_logged_in = false;
            sessionStorage.is_user_logged_in = false;
            $window.location.replace("/")
        });
    }

    //
    // Sidebar handlers
    //
    $scope.sidebar_init = function() {
        if (sessionStorage.sidebar_localhost == 'block')
            document.getElementById('sidebar-localhost').style.display = 'block';
    }

    $scope.open_localhost = function() {
        var localhost_sidebar_state = document.getElementById('sidebar-localhost').style.display;

        if (localhost_sidebar_state == '' || localhost_sidebar_state == 'none'){
            sessionStorage.sidebar_localhost = 'block';
            document.getElementById('sidebar-localhost').style.display = 'block';
        }
        else{
            sessionStorage.sidebar_localhost = '';
            document.getElementById('sidebar-localhost').style.display = 'none';
        }
    }
}
