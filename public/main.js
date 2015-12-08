$(document).ready(function(){
    onPageLoad(function() {
        $("#convertForm").submit(function (e) {
            e.preventDefault();
            if (!window.File || !window.FileReader || !window.FileList || !window.Blob) {
                alert('The File APIs are not fully supported in this browser.');
                return;
            }
            var form = $(this);
            var reader = new FileReader();
            reader.onload = function () {
                $.ajax({
                    url: form.attr("action"),
                    data: { "messages": window.atob(reader.result.split(",")[1])},
                    cache: false,
                    type: form.attr("method"),
                    success: function (data) {
                        switchPage("show", function () {
                            setTimeout(function(){
                                $("#whoami").html(data.whoami);
                            }, 20);
                            // bug?? have to wait 20 ms before setting or else resets, wat?
                        });
                    }
                });
            };
            reader.readAsDataURL(form.find("#file")[0].files[0]);
        });
    });
    startPureReplace();
});

