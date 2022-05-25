window.onload=function(){

    $.getJSON("./assets/words/"+ getTodayDate() + '.json', function( data ) {
        console.log(data);
        var items = [];
        $.each( data, function( key, val ) {
          items.push( "<li id='" + key + "' class='list-group-item'>" + val.word +":"+ val.def + "</li>" );
        });
       
        $( "<ul/>", {
            "class": "list-group",
            html: items.join( "" )
        }).appendTo( "#new-words" );
    });
}

function getTodayDate() {
    var today = new Date();
    var dd = String(today.getDate()).padStart(2, '0');
    var mm = String(today.getMonth() + 1).padStart(2, '0'); //January is 0!
    var yyyy = today.getFullYear();
    
    return yyyy+mm+dd;
}