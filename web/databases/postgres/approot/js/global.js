/* Navigation Start */
var xhr = new XMLHttpRequest();

xhr.open("GET", "/api/user/verify", true);
xhr.send();

xhr.onloadend = function(){
    if (xhr.status == 204) {
        jQuery('.navbar-nav').append('<li class="nav-item"><a class="nav-link" href="#">New Thread</a></li>');
        jQuery('.navbar-nav').append('<li class="nav-item"><a class="nav-link" href="#">Logout</a></li>');
    } else {
        jQuery('.navbar-nav').append('<li class="nav-item"><a class="nav-link" href="#">Login</a></li>');
        jQuery('.navbar-nav').append('<li class="nav-item"><a class="nav-link" href="#">Register</a></li>');
    }
};
/* Navigation End */