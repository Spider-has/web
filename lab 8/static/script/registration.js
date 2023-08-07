window.onload = () => {
    let form = document.getElementById("registration-form");

    if(form){
        form.addEventListener('submit', formSubmitAction);
    }
}

const emailInputId = "email-input";
const passwordInputId = "password-input";

const formSubmitAction = (event) => {
    event.preventDefault()
    console.log('отправка');
    checkFormValid(emailInputId, passwordInputId);
}

const checkFormValid = (emailInputId, passwordInputId) => {
    let emailInput = document.getElementById(emailInputId);
    let passwordInput = document.getElementById(passwordInputId);

    if(emailInput){
        console.log(emailInput.validity.valid);
    }
    
}