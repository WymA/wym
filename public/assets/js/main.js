window.onload = function () {
}

function getTodayDate() {
    var today = new Date();
    var dd = String(today.getDate()).padStart(2, '0');
    var mm = String(today.getMonth() + 1).padStart(2, '0'); //January is 0!
    var yyyy = today.getFullYear();

    return yyyy + mm + dd;
}

function getDailyNewWordJson(){
    $.getJSON("./assets/words/" + getTodayDate() + '.json', function (data) {
        console.log(data);
        var items = [];
        $.each(data, function (key, val) {
            var no = key+1;
            $('#new-words').append(`<div id="`+key+`" class="col-sm-4">
                                        <div class="card card-flip h-100">
                                            <div class="card-front text-white bg-dark">
                                                <div class="card-body">
                                                    <i class="fa fa-search fa-5x float-right"></i>
                                                    <h3 class="card-title">#`+ no + `</h3>
                                                    <p class="card-text">`+ val.word + `</p>
                                                </div>
                                            </div>
                                            <div class="card-back bg-white">
                                                <div class="card-body">
                                                    <h3 class="card-title">`+ val.word + `</h3>
                                                    <p class="card-text">`+ val.def + `</p>
                                                </div>
                                            </div>
                                        </div>
                                    </div>`);
        });
    });
}