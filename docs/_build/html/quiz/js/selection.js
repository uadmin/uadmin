$('.selection-group').unbind("click");
$('.selection-group').on("click",function(e){
    e.stopPropagation();
    var target = e.target;
    var parent_target = $(target).closest('.selection').find('input');
    if (parent_target && parent_target.length > 0){
        for (let index = 0; index < parent_target.length; index++) {
            $(parent_target[index]).removeAttr('checked');
        }
    }
    if ($(this).find('input').is(':enabled')){
        $(this).find('input').prop("checked", true);
        $(this).find('input').attr("checked", 'checked');    
    }    
});
