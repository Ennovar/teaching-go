var button = document.getElementsByClassName('_js_get-data-button')[0]

button.addEventListener('click', function(e){
  e.preventDefault();

  var xhr = new XMLHttpRequest();
  xhr.open('GET', 'api/data/get', true);
  xhr.send();

  xhr.onloadend = function() {
    if(xhr.status == 200) {
      alert(xhr.response);
    }
    else {
      if(xhr.response != undefined && xhr.response.length != 0) {
        alert('An HTTP ' + xhr.status + ' error has occurred: ' + xhr.status);
      }
      else {
        alert('An HTTP ' + xhr.status + ' error has occurred.');
      }
    }
  }
});
