var total = 0;
var skipped = 0;

$("#validate").click(function(){
    if(jQuery("#rQ1A2").is(":checked")){
        $("#Q1A1").css("color", "darkgreen");
        $("#Q1A2").css("color", "red");
        $("#Q1A3").css("color", "black");
        $("#Q1A4").css("color", "black");
    } else if(jQuery("#rQ1A3").is(":checked")){
        $("#Q1A1").css("color", "darkgreen");
        $("#Q1A2").css("color", "black");
        $("#Q1A3").css("color", "red");
        $("#Q1A4").css("color", "black");
    } else if(jQuery("#rQ1A4").is(":checked")){
        $("#Q1A1").css("color", "darkgreen");
        $("#Q1A2").css("color", "black");
        $("#Q1A3").css("color", "black");
        $("#Q1A4").css("color", "red");
    } else {
        $("#Q1A1").css("color", "darkgreen");
        $("#Q1A2").css("color", "black");
        $("#Q1A3").css("color", "black");
        $("#Q1A4").css("color", "black");
        if(jQuery("#rQ1A1").is(":checked")){
            total++;
        } else {
            skipped++;
        }
    }

    if(jQuery("#rQ2A1").is(":checked")){
        $("#Q2A1").css("color", "red");
        $("#Q2A2").css("color", "black");
        $("#Q2A3").css("color", "darkgreen");
        $("#Q2A4").css("color", "black");
    } else if(jQuery("#rQ2A2").is(":checked")){
        $("#Q2A1").css("color", "black");
        $("#Q2A2").css("color", "red");
        $("#Q2A3").css("color", "darkgreen");
        $("#Q2A4").css("color", "black");
    } else if(jQuery("#rQ2A4").is(":checked")){
        $("#Q2A1").css("color", "black");
        $("#Q2A2").css("color", "black");
        $("#Q2A3").css("color", "darkgreen");
        $("#Q2A4").css("color", "red");
    } else {
        $("#Q2A1").css("color", "black");
        $("#Q2A2").css("color", "black");
        $("#Q2A3").css("color", "darkgreen");
        $("#Q2A4").css("color", "black");
        if(jQuery("#rQ2A3").is(":checked")){
            total++;
        } else {
            skipped++;
        }
    }

    if(jQuery("#rQ3A1").is(":checked")){
        $("#Q3A1").css("color", "red");
        $("#Q3A2").css("color", "darkgreen");
        $("#Q3A3").css("color", "black");
        $("#Q3A4").css("color", "black");
    } else if(jQuery("#rQ3A3").is(":checked")){
        $("#Q3A1").css("color", "black");
        $("#Q3A2").css("color", "darkgreen");
        $("#Q3A3").css("color", "red");
        $("#Q3A4").css("color", "black");
    } else if(jQuery("#rQ3A4").is(":checked")){
        $("#Q3A1").css("color", "black");
        $("#Q3A2").css("color", "darkgreen");
        $("#Q3A3").css("color", "black");
        $("#Q3A4").css("color", "red");
    } else {
        $("#Q3A1").css("color", "black");
        $("#Q3A2").css("color", "darkgreen");
        $("#Q3A3").css("color", "black");
        $("#Q3A4").css("color", "black");
        if(jQuery("#rQ3A2").is(":checked")){
            total++;
        } else {
            skipped++;
        }
    }

    if(jQuery("#rQ4A2").is(":checked")){
        $("#Q4A1").css("color", "darkgreen");
        $("#Q4A2").css("color", "red");
        $("#Q4A3").css("color", "black");
        $("#Q4A4").css("color", "black");
    } else if(jQuery("#rQ4A3").is(":checked")){
        $("#Q4A1").css("color", "darkgreen");
        $("#Q4A2").css("color", "black");
        $("#Q4A3").css("color", "red");
        $("#Q4A4").css("color", "black");
    } else if(jQuery("#rQ4A4").is(":checked")){
        $("#Q4A1").css("color", "darkgreen");
        $("#Q4A2").css("color", "black");
        $("#Q4A3").css("color", "black");
        $("#Q4A4").css("color", "red");
    } else {
        $("#Q4A1").css("color", "darkgreen");
        $("#Q4A2").css("color", "black");
        $("#Q4A3").css("color", "black");
        $("#Q4A4").css("color", "black");
        if(jQuery("#rQ4A1").is(":checked")){
            total++;
        } else {
            skipped++;
        }
    }

    if(jQuery("#rQ5A1").is(":checked")){
        $("#Q5A1").css("color", "red");
        $("#Q5A2").css("color", "darkgreen");
        $("#Q5A3").css("color", "black");
        $("#Q5A4").css("color", "black");
    } else if(jQuery("#rQ5A3").is(":checked")){
        $("#Q5A1").css("color", "black");
        $("#Q5A2").css("color", "darkgreen");
        $("#Q5A3").css("color", "red");
        $("#Q5A4").css("color", "black");
    } else if(jQuery("#rQ5A4").is(":checked")){
        $("#Q5A1").css("color", "black");
        $("#Q5A2").css("color", "darkgreen");
        $("#Q5A3").css("color", "black");
        $("#Q5A4").css("color", "red");
    } else {
        $("#Q5A1").css("color", "black");
        $("#Q5A2").css("color", "darkgreen");
        $("#Q5A3").css("color", "black");
        $("#Q5A4").css("color", "black");
        if(jQuery("#rQ5A2").is(":checked")){
            total++;
        } else {
            skipped++;
        }
    }

    if(jQuery("#rQ6A1").is(":checked")){
        $("#Q6A1").css("color", "red");
        $("#Q6A2").css("color", "black");
        $("#Q6A3").css("color", "black");
        $("#Q6A4").css("color", "darkgreen");
    } else if(jQuery("#rQ6A2").is(":checked")){
        $("#Q6A1").css("color", "black");
        $("#Q6A2").css("color", "red");
        $("#Q6A3").css("color", "black");
        $("#Q6A4").css("color", "darkgreen");
    } else if(jQuery("#rQ6A3").is(":checked")){
        $("#Q6A1").css("color", "black");
        $("#Q6A2").css("color", "black");
        $("#Q6A3").css("color", "red");
        $("#Q6A4").css("color", "darkgreen");
    } else {
        $("#Q6A1").css("color", "black");
        $("#Q6A2").css("color", "black");
        $("#Q6A3").css("color", "black");
        $("#Q6A4").css("color", "darkgreen");
        if(jQuery("#rQ6A4").is(":checked")){
            total++;
        } else {
            skipped++;
        }
    }

    if(jQuery("#rQ7A1").is(":checked")){
        $("#Q7A1").css("color", "red");
        $("#Q7A2").css("color", "black");
        $("#Q7A3").css("color", "black");
        $("#Q7A4").css("color", "darkgreen");
    } else if(jQuery("#rQ7A2").is(":checked")){
        $("#Q7A1").css("color", "black");
        $("#Q7A2").css("color", "red");
        $("#Q7A3").css("color", "black");
        $("#Q7A4").css("color", "darkgreen");
    } else if(jQuery("#rQ7A3").is(":checked")){
        $("#Q7A1").css("color", "black");
        $("#Q7A2").css("color", "black");
        $("#Q7A3").css("color", "red");
        $("#Q7A4").css("color", "darkgreen");
    } else {
        $("#Q7A1").css("color", "black");
        $("#Q7A2").css("color", "black");
        $("#Q7A3").css("color", "black");
        $("#Q7A4").css("color", "darkgreen");
        if(jQuery("#rQ7A4").is(":checked")){
            total++;
        } else {
            skipped++;
        }
    }

    if(jQuery("#rQ8A1").is(":checked")){
        $("#Q8A1").css("color", "red");
        $("#Q8A2").css("color", "black");
        $("#Q8A3").css("color", "darkgreen");
        $("#Q8A4").css("color", "black");
    } else if(jQuery("#rQ8A2").is(":checked")){
        $("#Q8A1").css("color", "black");
        $("#Q8A2").css("color", "red");
        $("#Q8A3").css("color", "darkgreen");
        $("#Q8A4").css("color", "black");
    } else if(jQuery("#rQ8A4").is(":checked")){
        $("#Q8A1").css("color", "black");
        $("#Q8A2").css("color", "black");
        $("#Q8A3").css("color", "darkgreen");
        $("#Q8A4").css("color", "red");
    } else {
        $("#Q8A1").css("color", "black");
        $("#Q8A2").css("color", "black");
        $("#Q8A3").css("color", "darkgreen");
        $("#Q8A4").css("color", "black");
        if(jQuery("#rQ8A3").is(":checked")){
            total++;
        } else {
            skipped++;
        }
    }

    if(jQuery("#rQ9A1").is(":checked")){
        $("#Q9A1").css("color", "red");
        $("#Q9A2").css("color", "black");
        $("#Q9A3").css("color", "darkgreen");
        $("#Q9A4").css("color", "black");
    } else if(jQuery("#rQ9A2").is(":checked")){
        $("#Q9A1").css("color", "black");
        $("#Q9A2").css("color", "red");
        $("#Q9A3").css("color", "darkgreen");
        $("#Q9A4").css("color", "black");
    } else if(jQuery("#rQ9A4").is(":checked")){
        $("#Q9A1").css("color", "black");
        $("#Q9A2").css("color", "black");
        $("#Q9A3").css("color", "darkgreen");
        $("#Q9A4").css("color", "red");
    } else {
        $("#Q9A1").css("color", "black");
        $("#Q9A2").css("color", "black");
        $("#Q9A3").css("color", "darkgreen");
        $("#Q9A4").css("color", "black");
        if(jQuery("#rQ9A3").is(":checked")){
            total++;
        } else {
            skipped++;
        }
    }

    if(jQuery("#rQ10A2").is(":checked")){
        $("#Q10A1").css("color", "darkgreen");
        $("#Q10A2").css("color", "red");
        $("#Q10A3").css("color", "black");
        $("#Q10A4").css("color", "black");
    } else if(jQuery("#rQ10A3").is(":checked")){
        $("#Q10A1").css("color", "darkgreen");
        $("#Q10A2").css("color", "black");
        $("#Q10A3").css("color", "red");
        $("#Q10A4").css("color", "black");
    } else if(jQuery("#rQ10A4").is(":checked")){
        $("#Q10A1").css("color", "darkgreen");
        $("#Q10A2").css("color", "black");
        $("#Q10A3").css("color", "black");
        $("#Q10A4").css("color", "red");
    } else {
        $("#Q10A1").css("color", "darkgreen");
        $("#Q10A2").css("color", "black");
        $("#Q10A3").css("color", "black");
        $("#Q10A4").css("color", "black");
        if(jQuery("#rQ10A1").is(":checked")){
            total++;
        } else {
            skipped++;
        }
    }

    if(jQuery("#rQ11A1").is(":checked")){
        $("#Q11A1").css("color", "red");
        $("#Q11A2").css("color", "darkgreen");
        $("#Q11A3").css("color", "black");
        $("#Q11A4").css("color", "black");
    } else if(jQuery("#rQ11A3").is(":checked")){
        $("#Q11A1").css("color", "black");
        $("#Q11A2").css("color", "darkgreen");
        $("#Q11A3").css("color", "red");
        $("#Q11A4").css("color", "black");
    } else if(jQuery("#rQ11A4").is(":checked")){
        $("#Q11A1").css("color", "black");
        $("#Q11A2").css("color", "darkgreen");
        $("#Q11A3").css("color", "black");
        $("#Q11A4").css("color", "red");
    } else {
        $("#Q11A1").css("color", "black");
        $("#Q11A2").css("color", "darkgreen");
        $("#Q11A3").css("color", "black");
        $("#Q11A4").css("color", "black");
        if(jQuery("#rQ11A2").is(":checked")){
            total++;
        } else {
            skipped++;
        }
    }

    if(jQuery("#rQ12A1").is(":checked")){
        $("#Q12A1").css("color", "red");
        $("#Q12A2").css("color", "black");
        $("#Q12A3").css("color", "darkgreen");
        $("#Q12A4").css("color", "black");
    } else if(jQuery("#rQ12A2").is(":checked")){
        $("#Q12A1").css("color", "black");
        $("#Q12A2").css("color", "red");
        $("#Q12A3").css("color", "darkgreen");
        $("#Q12A4").css("color", "black");
    } else if(jQuery("#rQ12A4").is(":checked")){
        $("#Q12A1").css("color", "black");
        $("#Q12A2").css("color", "black");
        $("#Q12A3").css("color", "darkgreen");
        $("#Q12A4").css("color", "red");
    } else {
        $("#Q12A1").css("color", "black");
        $("#Q12A2").css("color", "black");
        $("#Q12A3").css("color", "darkgreen");
        $("#Q12A4").css("color", "black");
        if(jQuery("#rQ12A3").is(":checked")){
            total++;
        } else {
            skipped++;
        }
    }

    if(jQuery("#rQ13A1").is(":checked")){
        $("#Q13A1").css("color", "red");
        $("#Q13A2").css("color", "black");
        $("#Q13A3").css("color", "black");
        $("#Q13A4").css("color", "darkgreen");
    } else if(jQuery("#rQ13A2").is(":checked")){
        $("#Q13A1").css("color", "black");
        $("#Q13A2").css("color", "red");
        $("#Q13A3").css("color", "black");
        $("#Q13A4").css("color", "darkgreen");
    } else if(jQuery("#rQ13A3").is(":checked")){
        $("#Q13A1").css("color", "black");
        $("#Q13A2").css("color", "black");
        $("#Q13A3").css("color", "red");
        $("#Q13A4").css("color", "darkgreen");
    } else {
        $("#Q13A1").css("color", "black");
        $("#Q13A2").css("color", "black");
        $("#Q13A3").css("color", "black");
        $("#Q13A4").css("color", "darkgreen");
        if(jQuery("#rQ13A4").is(":checked")){
            total++;
        } else {
            skipped++;
        }
    }

    if(jQuery("#rQ14A1").is(":checked")){
        $("#Q14A1").css("color", "red");
        $("#Q14A2").css("color", "darkgreen");
        $("#Q14A3").css("color", "black");
        $("#Q14A4").css("color", "black");
    } else if(jQuery("#rQ14A3").is(":checked")){
        $("#Q14A1").css("color", "black");
        $("#Q14A2").css("color", "darkgreen");
        $("#Q14A3").css("color", "red");
        $("#Q14A4").css("color", "black");
    } else if(jQuery("#rQ14A4").is(":checked")){
        $("#Q14A1").css("color", "black");
        $("#Q14A2").css("color", "darkgreen");
        $("#Q14A3").css("color", "black");
        $("#Q14A4").css("color", "red");
    } else {
        $("#Q14A1").css("color", "black");
        $("#Q14A2").css("color", "darkgreen");
        $("#Q14A3").css("color", "black");
        $("#Q14A4").css("color", "black");
        if(jQuery("#rQ14A2").is(":checked")){
            total++;
        } else {
            skipped++;
        }
    }

    if(jQuery("#rQ15A1").is(":checked")){
        $("#Q15A1").css("color", "red");
        $("#Q15A2").css("color", "black");
        $("#Q15A3").css("color", "darkgreen");
        $("#Q15A4").css("color", "black");
    } else if(jQuery("#rQ15A2").is(":checked")){
        $("#Q15A1").css("color", "black");
        $("#Q15A2").css("color", "red");
        $("#Q15A3").css("color", "darkgreen");
        $("#Q15A4").css("color", "black");
    } else if(jQuery("#rQ15A4").is(":checked")){
        $("#Q15A1").css("color", "black");
        $("#Q15A2").css("color", "black");
        $("#Q15A3").css("color", "darkgreen");
        $("#Q15A4").css("color", "red");
    } else {
        $("#Q15A1").css("color", "black");
        $("#Q15A2").css("color", "black");
        $("#Q15A3").css("color", "darkgreen");
        $("#Q15A4").css("color", "black");
        if(jQuery("#rQ15A3").is(":checked")){
            total++;
        } else {
            skipped++;
        }
    }

    // Show the results
    $("#result").show();
    $("#skipped").show();

    // Prints the results
    $("#result").html("Your total score is <b>" + total + "</b> out of 15.");
    if (skipped != 0) {
        $("#skipped").html("You skipped <b>" + skipped + "</b> questions.");
    }

    // Disable the submit button
    $("#validate").attr("disabled", "disabled");

    // Disable the radio buttons
    $("#rQ1A1").attr("disabled", "disabled");
    $("#rQ1A2").attr("disabled", "disabled");
    $("#rQ1A3").attr("disabled", "disabled");
    $("#rQ1A4").attr("disabled", "disabled");
    $("#rQ2A1").attr("disabled", "disabled");
    $("#rQ2A2").attr("disabled", "disabled");
    $("#rQ2A3").attr("disabled", "disabled");
    $("#rQ2A4").attr("disabled", "disabled");
    $("#rQ3A1").attr("disabled", "disabled");
    $("#rQ3A2").attr("disabled", "disabled");
    $("#rQ3A3").attr("disabled", "disabled");
    $("#rQ3A4").attr("disabled", "disabled");
    $("#rQ4A1").attr("disabled", "disabled");
    $("#rQ4A2").attr("disabled", "disabled");
    $("#rQ4A3").attr("disabled", "disabled");
    $("#rQ4A4").attr("disabled", "disabled");
    $("#rQ5A1").attr("disabled", "disabled");
    $("#rQ5A2").attr("disabled", "disabled");
    $("#rQ5A3").attr("disabled", "disabled");
    $("#rQ5A4").attr("disabled", "disabled");
    $("#rQ6A1").attr("disabled", "disabled");
    $("#rQ6A2").attr("disabled", "disabled");
    $("#rQ6A3").attr("disabled", "disabled");
    $("#rQ6A4").attr("disabled", "disabled");
    $("#rQ7A1").attr("disabled", "disabled");
    $("#rQ7A2").attr("disabled", "disabled");
    $("#rQ7A3").attr("disabled", "disabled");
    $("#rQ7A4").attr("disabled", "disabled");
    $("#rQ8A1").attr("disabled", "disabled");
    $("#rQ8A2").attr("disabled", "disabled");
    $("#rQ8A3").attr("disabled", "disabled");
    $("#rQ8A4").attr("disabled", "disabled");
    $("#rQ9A1").attr("disabled", "disabled");
    $("#rQ9A2").attr("disabled", "disabled");
    $("#rQ9A3").attr("disabled", "disabled");
    $("#rQ9A4").attr("disabled", "disabled");
    $("#rQ10A1").attr("disabled", "disabled");
    $("#rQ10A2").attr("disabled", "disabled");
    $("#rQ10A3").attr("disabled", "disabled");
    $("#rQ10A4").attr("disabled", "disabled");
    $("#rQ11A1").attr("disabled", "disabled");
    $("#rQ11A2").attr("disabled", "disabled");
    $("#rQ11A3").attr("disabled", "disabled");
    $("#rQ11A4").attr("disabled", "disabled");
    $("#rQ12A1").attr("disabled", "disabled");
    $("#rQ12A2").attr("disabled", "disabled");
    $("#rQ12A3").attr("disabled", "disabled");
    $("#rQ12A4").attr("disabled", "disabled");
    $("#rQ13A1").attr("disabled", "disabled");
    $("#rQ13A2").attr("disabled", "disabled");
    $("#rQ13A3").attr("disabled", "disabled");
    $("#rQ13A4").attr("disabled", "disabled");
    $("#rQ14A1").attr("disabled", "disabled");
    $("#rQ14A2").attr("disabled", "disabled");
    $("#rQ14A3").attr("disabled", "disabled");
    $("#rQ14A4").attr("disabled", "disabled");
    $("#rQ15A1").attr("disabled", "disabled");
    $("#rQ15A2").attr("disabled", "disabled");
    $("#rQ15A3").attr("disabled", "disabled");
    $("#rQ15A4").attr("disabled", "disabled");
});

$("#reset").click(function(){
    // Reset the total and skip values
    total = 0;
    skipped = 0;

    // Set the font color of the choices to default
    $("#Q1A1").css("color", "black");
    $("#Q1A2").css("color", "black");
    $("#Q1A3").css("color", "black");
    $("#Q1A4").css("color", "black");
    $("#Q2A1").css("color", "black");
    $("#Q2A2").css("color", "black");
    $("#Q2A3").css("color", "black");
    $("#Q2A4").css("color", "black");
    $("#Q3A1").css("color", "black");
    $("#Q3A2").css("color", "black");
    $("#Q3A3").css("color", "black");
    $("#Q3A4").css("color", "black");
    $("#Q4A1").css("color", "black");
    $("#Q4A2").css("color", "black");
    $("#Q4A3").css("color", "black");
    $("#Q4A4").css("color", "black");
    $("#Q5A1").css("color", "black");
    $("#Q5A2").css("color", "black");
    $("#Q5A3").css("color", "black");
    $("#Q5A4").css("color", "black");
    $("#Q6A1").css("color", "black");
    $("#Q6A2").css("color", "black");
    $("#Q6A3").css("color", "black");
    $("#Q6A4").css("color", "black");
    $("#Q7A1").css("color", "black");
    $("#Q7A2").css("color", "black");
    $("#Q7A3").css("color", "black");
    $("#Q7A4").css("color", "black");
    $("#Q8A1").css("color", "black");
    $("#Q8A2").css("color", "black");
    $("#Q8A3").css("color", "black");
    $("#Q8A4").css("color", "black");
    $("#Q9A1").css("color", "black");
    $("#Q9A2").css("color", "black");
    $("#Q9A3").css("color", "black");
    $("#Q9A4").css("color", "black");
    $("#Q10A1").css("color", "black");
    $("#Q10A2").css("color", "black");
    $("#Q10A3").css("color", "black");
    $("#Q10A4").css("color", "black");
    $("#Q11A1").css("color", "black");
    $("#Q11A2").css("color", "black");
    $("#Q11A3").css("color", "black");
    $("#Q11A4").css("color", "black");
    $("#Q12A1").css("color", "black");
    $("#Q12A2").css("color", "black");
    $("#Q12A3").css("color", "black");
    $("#Q12A4").css("color", "black");
    $("#Q13A1").css("color", "black");
    $("#Q13A2").css("color", "black");
    $("#Q13A3").css("color", "black");
    $("#Q13A4").css("color", "black");
    $("#Q14A1").css("color", "black");
    $("#Q14A2").css("color", "black");
    $("#Q14A3").css("color", "black");
    $("#Q14A4").css("color", "black");
    $("#Q15A1").css("color", "black");
    $("#Q15A2").css("color", "black");
    $("#Q15A3").css("color", "black");
    $("#Q15A4").css("color", "black");

    // Clear the checked radio buttons and enable them
    $("#rQ1A1").prop('checked', false).removeAttr("disabled");
    $("#rQ1A2").prop('checked', false).removeAttr("disabled");
    $("#rQ1A3").prop('checked', false).removeAttr("disabled");
    $("#rQ1A4").prop('checked', false).removeAttr("disabled");
    $("#rQ2A1").prop('checked', false).removeAttr("disabled");
    $("#rQ2A2").prop('checked', false).removeAttr("disabled");
    $("#rQ2A3").prop('checked', false).removeAttr("disabled");
    $("#rQ2A4").prop('checked', false).removeAttr("disabled");
    $("#rQ3A1").prop('checked', false).removeAttr("disabled");
    $("#rQ3A2").prop('checked', false).removeAttr("disabled");
    $("#rQ3A3").prop('checked', false).removeAttr("disabled");
    $("#rQ3A4").prop('checked', false).removeAttr("disabled");
    $("#rQ4A1").prop('checked', false).removeAttr("disabled");
    $("#rQ4A2").prop('checked', false).removeAttr("disabled");
    $("#rQ4A3").prop('checked', false).removeAttr("disabled");
    $("#rQ4A4").prop('checked', false).removeAttr("disabled");
    $("#rQ5A1").prop('checked', false).removeAttr("disabled");
    $("#rQ5A2").prop('checked', false).removeAttr("disabled");
    $("#rQ5A3").prop('checked', false).removeAttr("disabled");
    $("#rQ5A4").prop('checked', false).removeAttr("disabled");
    $("#rQ6A1").prop('checked', false).removeAttr("disabled");
    $("#rQ6A2").prop('checked', false).removeAttr("disabled");
    $("#rQ6A3").prop('checked', false).removeAttr("disabled");
    $("#rQ6A4").prop('checked', false).removeAttr("disabled");
    $("#rQ7A1").prop('checked', false).removeAttr("disabled");
    $("#rQ7A2").prop('checked', false).removeAttr("disabled");
    $("#rQ7A3").prop('checked', false).removeAttr("disabled");
    $("#rQ7A4").prop('checked', false).removeAttr("disabled");
    $("#rQ8A1").prop('checked', false).removeAttr("disabled");
    $("#rQ8A2").prop('checked', false).removeAttr("disabled");
    $("#rQ8A3").prop('checked', false).removeAttr("disabled");
    $("#rQ8A4").prop('checked', false).removeAttr("disabled");
    $("#rQ9A1").prop('checked', false).removeAttr("disabled");
    $("#rQ9A2").prop('checked', false).removeAttr("disabled");
    $("#rQ9A3").prop('checked', false).removeAttr("disabled");
    $("#rQ9A4").prop('checked', false).removeAttr("disabled");
    $("#rQ10A1").prop('checked', false).removeAttr("disabled");
    $("#rQ10A2").prop('checked', false).removeAttr("disabled");
    $("#rQ10A3").prop('checked', false).removeAttr("disabled");
    $("#rQ10A4").prop('checked', false).removeAttr("disabled");
    $("#rQ11A1").prop('checked', false).removeAttr("disabled");
    $("#rQ11A2").prop('checked', false).removeAttr("disabled");
    $("#rQ11A3").prop('checked', false).removeAttr("disabled");
    $("#rQ11A4").prop('checked', false).removeAttr("disabled");
    $("#rQ12A1").prop('checked', false).removeAttr("disabled");
    $("#rQ12A2").prop('checked', false).removeAttr("disabled");
    $("#rQ12A3").prop('checked', false).removeAttr("disabled");
    $("#rQ12A4").prop('checked', false).removeAttr("disabled");
    $("#rQ13A1").prop('checked', false).removeAttr("disabled");
    $("#rQ13A2").prop('checked', false).removeAttr("disabled");
    $("#rQ13A3").prop('checked', false).removeAttr("disabled");
    $("#rQ13A4").prop('checked', false).removeAttr("disabled");
    $("#rQ14A1").prop('checked', false).removeAttr("disabled");
    $("#rQ14A2").prop('checked', false).removeAttr("disabled");
    $("#rQ14A3").prop('checked', false).removeAttr("disabled");
    $("#rQ14A4").prop('checked', false).removeAttr("disabled");
    $("#rQ15A1").prop('checked', false).removeAttr("disabled");
    $("#rQ15A2").prop('checked', false).removeAttr("disabled");
    $("#rQ15A3").prop('checked', false).removeAttr("disabled");
    $("#rQ15A4").prop('checked', false).removeAttr("disabled");

    // Hide the results below the buttons
    $("#result").hide();
    $("#skipped").hide();

    // Enable the submit button
    $("#validate").removeAttr("disabled");
});