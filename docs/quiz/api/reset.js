$(".reset").click(function(){
    // Reset the total and skip values
    total = 0;
    skipped = 0;

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
    $("#validate").removeAttr("disabled");
});