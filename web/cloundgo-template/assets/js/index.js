$(document).ready(function() {
    $.ajax({
        url: "/js"
    }).then(function(data) {
       $('.greeting-time').append(data.now);
    });
});