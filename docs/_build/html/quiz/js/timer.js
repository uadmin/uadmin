var score = 0;
var skipped = 0;
var total = 0;

function validate(){
    // toggle correct color
    $('[data-correct="true"]').addClass('correct');
    var answers = $("body").find('input[checked]');
    if (answers && answers.length > 0){
        for (var i = 0; i < answers.length; i++){
            var selected_item = $(answers[i]).closest('.selection-group');
            if (selected_item && selected_item.length > 0){
                if (selected_item.attr('data-correct') == "true"){
                    score++;
                }else{
                    selected_item.addClass('incorrect');
                }
                total++;
            }
        }
    }
    var skipped_questions_parent = $("body").find('.selection');
    if (skipped_questions_parent && skipped_questions_parent.length > 0){
        for (var i = 0; i < skipped_questions_parent.length; i++) {
            if ($(skipped_questions_parent[i]) && $(skipped_questions_parent[i]).length > 0){
                var skipped_items = $(skipped_questions_parent[i]).find("input[checked]");
                if (skipped_items.length == 0){
                    skipped++;
                    total++;
                }
            }            
        }
    }

    // Show the results
    $("#result").show();

    // Prints the results
    $("#result").html("<br><hr>Your total score is <b>" + score + "</b> out of " + total + ".");
    if (skipped != 0) {
        $("#skipped").show();
        $("#skipped").html("You skipped <b>" + skipped + "</b> questions.");
    }
    console.log(skipped);

    // Disable the submit button
    $(".validate").attr("disabled", "disabled");

    // Disable the radio buttons
    var skipped_questions_parent = $("body").find('.selection');
    if (skipped_questions_parent && skipped_questions_parent.length > 0){
        for (var i = 0; i < skipped_questions_parent.length; i++) {
            if ($(skipped_questions_parent[i]) && $(skipped_questions_parent[i]).length > 0){
                var allitems = $(skipped_questions_parent[i]).find("input");
                if (allitems && allitems.length > 0){
                    for (var index = 0; index < allitems.length; index++) {
                        $(allitems[index])[0].disabled = true;
                    }
                }
            }            
        }
    }
}

function resetFunc(){
    // Reset the score, total and skip values
    score = 0;
    skipped = 0;
    total = 0;

    var skipped_questions_parent = $("body").find('.selection');
    if (skipped_questions_parent && skipped_questions_parent.length > 0){
        for (var i = 0; i < skipped_questions_parent.length; i++) {
            if ($(skipped_questions_parent[i]) && $(skipped_questions_parent[i]).length > 0){
                var allitems = $(skipped_questions_parent[i]).find("input");
                if (allitems && allitems.length > 0){
                    for (var index = 0; index < allitems.length; index++) {
                        $(allitems[index]).removeAttr('disabled');
                        $(allitems[index]).removeAttr('checked');
                        $(allitems[index]).prop('checked',false);
                        $(allitems[index]).closest('.selection-group').removeClass('incorrect').removeClass('correct');
                    }
                }
            }            
        }
    }

    // Hide the results below the buttons
    $("#result").hide();
    $("#skipped").hide();

    // Enable the submit button
    $(".validate").removeAttr("disabled");
}

function loadTimer(){
    jQuery(function ($) {
        var fiveMinutes = 60 * 9.99,
            display = $('#countdown').html("TIME LEFT:&nbsp;<b>10:00</b>");
        startTimer(fiveMinutes, display);
    });
}

function startTimer(duration, display) {
    var timer = duration, minutes, seconds;
    var tt = setInterval(function () {
        minutes = parseInt(timer / 60, 10);
        seconds = parseInt(timer % 60, 10);

        minutes = minutes < 10 ? "0" + minutes : minutes;
        seconds = seconds < 10 ? "0" + seconds : seconds;

        display.html("TIME LEFT:&nbsp;<b>" + minutes + ":" + seconds + "</b>");

        if (timer > 0){
            --timer;
        } else {
            alert("TIME'S UP!");
            clearInterval(tt);
            validate();
        }
    }, 1000);
    $('.validate').unbind("click");
    $('.reset').unbind("click");
    $('.validate').on("click",function(){    
        clearInterval(tt);
        alert("Submitted. Scroll down below for the results.");
        validate();
    });
    $('.reset').on("click",function(){
        clearInterval(tt);
        resetFunc();
        loadTimer();
    });
}