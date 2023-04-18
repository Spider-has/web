window.onload = () => {
    const body = document.createElement("div");
    if (body){
        console.log("success")
    }
    body.innerHTML = '<Anton>  ANTON  </Anton>';
    
    const div = document.getElementById("org_div1");
    document.body.insertBefore(body, div);
}


console.log('aaaa');