function validDate() {
    let startDate = document.getElementById("startDate").value;
    let endDate = document.getElementById("endDate").value;

    let d1 = new Date(startDate);
    let d2 = new Date(endDate);

    if (d1 > d2) {
        alert("Invalid start & end date");
        return false;
    }
    return true;
}