let originalImage = ""

function validDate() {
    let startDate = document.getElementById("startDate").value;
    let endDate = document.getElementById("endDate").value;

    let d1 = new Date(startDate);
    let d2 = new Date(endDate);

    if (d1 > d2) {
        alert("Invalid start & end date.");
        return false;
    }
    return true;
}

function updatePreview() {
    if (originalImage == "") {
        // originalImage = document.getElementById("image-preview").src;
        originalImage = "/public/images/nopreview.png";
    }

    let img = ""
    let image = document.getElementById("inputImage").files;
    if (image.length == 0) img = originalImage;
    else img = URL.createObjectURL(image[0]);
    document.getElementById("image-preview").src = img;
}