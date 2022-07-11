function onSearchSubmit() {
    let searchTerm = document.getElementById("searchTerm").value;
    window.location = "/search/" + searchTerm;
    return false;
}

function TrackEntry(id, type) {
    console.log("TrackEntry: " + id + " " + type);
    window.location = "/track/" + type + "/" + id;
}

function UntrackEntry(id, type) {
    console.log("UntrackEntry: " + id + " " + type);
    window.location = "/untrack/" + type + "/" + id;
}

function showHideBooks(id) {
    let bookList = document.getElementById(id);
    if (bookList.style.display == "block") {
        bookList.style.display = "none";
    } else {
        bookList.style.display = "block";
    }
}