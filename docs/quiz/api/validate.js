var total = 0;
var skipped = 0;

$(".validate").click(function(){
    // toggle correct color
    $('[data-correct="true"]').addClass('correct');
    var answers = $("body").find('input[checked]');
    console.log(answers.length);
    if (answers && answers.length > 0){
        for (var i = 0; i < answers.length; i++){
            var selected_item = $(answers[i]).closest('.selection-group');
            if (selected_item && selected_item.length > 0){
                if (selected_item.attr('data-correct') == "true"){
                    total++;
                }else{
                    selected_item.addClass('incorrect');
                }
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
                }
            }            
        }
    }

    // Show the results
    $("#result").show();
    $("#skipped").show();

    // Prints the results
    $("#result").html("<br><hr>Your total score is <b>" + total + "</b> out of 15.");
    if (skipped != 0) {
        $("#skipped").html("You skipped <b>" + skipped + "</b> questions.");
    }

    // Disable the submit button
    $("#validate").attr("disabled", "disabled");

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
});



