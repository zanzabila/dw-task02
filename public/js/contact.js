function submitData() {
    let name = document.getElementById("name").value;
    let email = document.getElementById("email").value;
    let phone = document.getElementById("phone").value;
    let subject = document.getElementById("subject").value;
    let message = document.getElementById("message").value;

    if (name == "") { return alert("Nama harus diisi!"); }
    else if (email == "") { return alert("Email harus diisi!"); }
    else if (phone == "") { return alert("Phone harus diisi!"); }
    else if (subject == "") { return alert("Subject harus dipilih!"); }
    else if (message == "") { return alert("Message harus diisi!"); }

    let emailReceiver = "zanzabila.rayhan@gmail.com"
    let a = document.createElement('a');
    a.href = `mailto:${emailReceiver}?subject=${subject}&body=Halo, nama saya ${name}, ${message}. Silakan kontak saya di nomor ${phone}, terima kasih.`;
    a.click();
}