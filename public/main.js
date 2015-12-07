$(document).ready(function(){

    $("#convertForm").submit(function(e){
        e.preventDefault();
        if (!window.File || !window.FileReader || !window.FileList || !window.Blob) {
            alert('The File APIs are not fully supported in this browser.');
            return;
        }
        file = $("#file")[0].files[0];
        fr = new FileReader();
        fr.onload = function(){
            $.ajax({
                url: "/convert",
                data: {"messages": fr.result.toString()},
                cache: false,
                contentType: 'multipart/form-data',
                processData: false,
                type: 'POST',
                success: function(data){
                    page("show", function() {
                        $("#whoami").val(data.whoami);
                    });
                }
            });
        };
        fr.readAsDataURL(file);

    });
});

