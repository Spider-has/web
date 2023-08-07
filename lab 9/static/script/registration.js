

const emailInputId = "email-input";
const passwordInputId = "password-input";
const formErrorMessage = "form-error-message";
const emailErrorMessageId = "email-invalid-message";
const passwordErrorMessageId = "password-invalid-message";
const passwordButtonId = "show-password-button";
const formErrorTextId = "form-error-text";

window.onload = () => {
    let form = document.getElementById("registration-form");
    let passwordButton = document.getElementById(passwordButtonId);

    if (passwordButton) {
        let shown = false;
        passwordButton.addEventListener('click', () => {
            changePasswordVisible(passwordButtonId, passwordInputId, shown);
            shown = !shown;
        })
    }
    if (form) {
        form.addEventListener('submit', formSubmitAction);
    }
}

const changePasswordVisible = (passwordIconId, passwordInputId, shown) => {
    let passwordInput = document.getElementById(passwordInputId);
    let passwordIcon = document.getElementById(passwordIconId);
    if (passwordInput) {
        if (!shown) {
            passwordInput.setAttribute('type', 'text');
        }
        else {
            passwordInput.setAttribute('type', 'password');
        }
    }

    if (passwordIcon) {
        if (!shown) {
            passwordIcon.src = "../static/img/admin/eye-off.svg";
        }
        else {
            passwordIcon.src = "../static/img/admin/eye.svg";
        }
    }

}

const url = "/api/login"

const formSubmitAction = (event) => {
    event.preventDefault()
   
    if(checkFormValid(emailInputId, passwordInputId))
    {
        doFetch(url, getAllInputsData());
    }
}

const postCreateUrl = "http://localhost:3000/admin/post-settings"

async function doFetch(url, data) {
    const response = await fetch(url,
        {
            method: 'POST', // Здесь так же могут быть GET, PUT, DELETE
            headers: { 'Content-Type': 'multipart/form-data' },
            body: JSON.stringify(data), // Тело запроса в JSON-формате
        }).then((response) => {
            if (!response.ok) {
                throw new Error('Error occurred!')
            }
            console.log('победа');
            window.location.assign(postCreateUrl);
            console.log('рок');
        }).catch((e) => {
            console.log("asfasffas");
            invalidFormData(emailInputId, passwordInputId)
        });
}

const getAllInputsData = () => {
    let inputsData = {};

    inputsData.email = getInputsTextData(emailInputId);
    inputsData.password = getInputsTextData(passwordInputId);

    return inputsData
}

const getInputsTextData = (inputId) => {
    let textInput = document.getElementById(inputId);
    if (textInput) {
        return textInput.value;
    }
    else {
        return "";
    }
}

const invalidFormData = (emailInputId, passwordInputId) => {
    let emailInput = document.getElementById(emailInputId);
    let passwordInput = document.getElementById(passwordInputId);
    let FormErrorMessage = document.getElementById(formErrorMessage);
    let FormErrorText = document.getElementById(formErrorTextId);

    FormErrorText.innerHTML = "Email or password is incorrect.";
    emailInput.classList.add('admin-form-input_invalid');
    FormErrorMessage.classList.add("admin-error-message_shown");
    passwordInput.classList.add('admin-form-input_invalid');
}

const checkFormValid = (emailInputId, passwordInputId) => {
    let emailInput = document.getElementById(emailInputId);
    let passwordInput = document.getElementById(passwordInputId);
    let FormErrorMessage = document.getElementById(formErrorMessage);
    let emailErrorMessage = document.getElementById(emailErrorMessageId);
    let passwordErrorMessage = document.getElementById(passwordErrorMessageId);

    if (emailInput && passwordInput) {
        if (!emailInput.validity.valid || passwordInput.value === "") {
            if (!emailInput.validity.valid) {
                emailInput.classList.add('admin-form-input_invalid');
                if (FormErrorMessage) {
                    FormErrorMessage.classList.add("admin-error-message_shown");
                }
                if (emailErrorMessage) {
                    emailErrorMessage.classList.add('invalid-message_shown');
                    if (emailInput.value === "") {
                        emailErrorMessage.innerHTML = "Email is required."
                    }
                    else {
                        emailErrorMessage.innerHTML = "Incorrect email format. Correct format is ****@**.***"
                    }
                }
            }
            else {
                emailInput.classList.remove('admin-form-input_invalid');
                emailErrorMessage.classList.remove('invalid-message_shown');
            }
            if (passwordInput.value === "") {
                if (FormErrorMessage) {
                    FormErrorMessage.classList.add("admin-error-message_shown");
                }
                passwordInput.classList.add('admin-form-input_invalid');
                if (FormErrorMessage) {
                    FormErrorMessage.classList.add("admin-error-message_shown");
                }
                if (passwordInput.value === "" && passwordErrorMessage) {
                    passwordErrorMessage.classList.add('invalid-message_shown');
                }
            }
            else {
                passwordInput.classList.remove('admin-form-input_invalid');
                passwordErrorMessage.classList.remove('invalid-message_shown');
            }
            return false
        }
        else {
            if (FormErrorMessage) {
                FormErrorMessage.classList.remove("admin-error-message_shown");

            }
            passwordInput.classList.remove('admin-form-input_invalid');
            passwordErrorMessage.classList.remove('invalid-message_shown');
            emailInput.classList.remove('admin-form-input_invalid');
            emailErrorMessage.classList.remove('invalid-message_shown');
            return true
        }
    }
    return false
}