(function(){
	"using strict";

	function setType(t) {
		if (t == "2") {
      $("select[name='ModelName']").closest("tr").show();
      $("select[name='ModelName']").select2({allowClear:true});
      $("select[name='Field']").closest("tr").show();
      $("select[name='Field']").select2({allowClear:true});
      $("input[name='PrimaryKey']").closest("tr").show();
      $("input[name='StaticPath']").closest("tr").hide();
    } else if (t == "1") {
      $("select[name='ModelName']").closest("tr").hide();
      $("select[name='Field']").closest("tr").hide();
      $("input[name='PrimaryKey']").closest("tr").hide();
      $("input[name='StaticPath']").closest("tr").show();
    }
	}

	// Show fields depending of the test type
	$("select[name='Type']").on("select2:select", function (e) {
		setType(e.params.data.id);
	});
	$(document).ready(function(){
		setType($('select[name="Type"]').val());
	});

	// Load Models
	/*
	$.get("/api/get_models/", function(data){
		$.each(data, function(k,v){
			$("select[name='ModelName']").append(new Option(v, k, false, false));
		});
		$("select[name='ModelName']").trigger("change");
	}, "json");
	*/

	// Load fields
	$("select[name='ModelName']").on("select2:select", function (e) {
		$.get("/api/get_fields/", {m:$('select[name="ModelName"]').select2('data')[0].text}, function(data){

			$("select[name='Field']").val(null).trigger('change');
			 $("select[name='Field']").html("").trigger('change');
	    $.each(data, function(k,v){
  	    $("select[name='Field']").append(new Option(v, k, false, false));
	    });
  	  $("select[name='Field']").trigger("change");
	  }, "json");
	});

	// Hide all fields
	$("input[name='StaticPath']").closest("tr").hide();
	$("select[name='ModelName']").closest("tr").hide();
	$("select[name='Field']").closest("tr").hide();
	$("input[name='PrimaryKey']").closest("tr").hide();
})();
