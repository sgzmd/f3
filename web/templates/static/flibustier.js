function onSearchSubmit() {
    let searchTerm = document.getElementById("searchTerm").value;
    window.location = "/search/" + searchTerm;
    return false;
}

function TrackUntrackEntry(id, type) {
    console.log("TrackUntrackEntry: " + id + " " + type);
    window.location = "/track/" + type + "/" + id;
}
