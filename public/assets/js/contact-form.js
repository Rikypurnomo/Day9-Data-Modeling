function getData(){
    
    let name = document.getElementById("name").value
    let email = document.getElementById("email").value
    let phoneNumber = document.getElementById("phone").value
    let subject = document.getElementById("subject").value
    let message = document.getElementById("message").value

    if(name =="") {
        return alert("Nama tidak boleh kosong")
    } else if(email == "") {
        return alert("email tidak boleh kosong")
    } else if(phoneNumber == ""){
        return alert("telepon tidak boleh kosong")
    } else if(subject == ""){
        return alert("subject tidak boleh kosong")
    } else if (message == ""){
        return alert("pesan tidak boleh kosong")
    }

    let emailReciver = "rikilah930@gmail.com"

    let mailTo = document.createElement('a')
    mailTo.href = `mailTo:${emailReciver}?subject=${subject}&body=Halo nama saya ${name}, ${message}, nomor telepon saya ${phoneNumber}`
    mailTo.click()

    let users = {
        myName: name,
        myEmail: email,
        myPhone: phoneNumber,
        mySubject: subject,
        myMessage: message
    }
}
