function onSearchSubmit() {
    let searchTerm = document.getElementById("searchTerm").value;
    window.location = "/search/" + searchTerm;
    return false;
}


// Generates onSubmit for HTML form
