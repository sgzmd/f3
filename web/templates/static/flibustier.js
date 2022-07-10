function onSearchSubmit() {
    let searchTerm = document.getElementById("searchTerm").value;
    window.location = "/search/" + searchTerm;
}
