let authorAvatarData = {
    base64: undefined,
}

let heroImg1Data = {
    base64: undefined,
}

let heroImg2Data = {
    base64: undefined,
}

const formId = "artcileForm";
const formTitleId = "formPostTitle";
const formSubtitleId = "formPostSubtitle";
const formAuthorNameId = "formAuthorName";
const formAuthorAvatarId = "formAuthorAvatar";
const formPublishDateId = "formPublishDate";
const formHeroImg1Id = "hero-img-1-input";
const formHeroImg2Id = "hero-img-2-input";
const formContentId = "formContent";

const AuthorAvatarImage = "AuthorAvatarImg";
const AuthorAvatarImagePreviewDiv = "image-input-preview";
const defaultAuthorAvatarText = "onload-author-avatar-text-area";
const AuthorAvatarSettings =  "author-avatar-settings";
const removeAuthorAvatar = "remove-Author-Avatar-button";

const heroImg1PreviewDiv = "hero-img-1-preview-div";
const heroImg1Img = "hero-img-1-preview";
const defaultHeroImg1Text = "hero-img-1-upload-text";
const HeroImg1Settings = "hero-img-1-settings";
const removeHeroImg1 = "remove-hero-img-1-button";
const HeroImg1Description = "hero-img-1-description";

const heroImg2PreviewDiv = "hero-img-2-preview-div";
const heroImg2Img = "hero-img-2-preview";
const defaultHeroImg2Text = "hero-img-2-upload-text";
const HeroImg2Settings = "hero-img-2-settings";
const removeHeroImg2 = "remove-hero-img-2-button";
const HeroImg2Description = "hero-img-2-description";


const acticlePreviewImg = "arcticle-preview-image";
const postCardPreviewImg = "post-card-preview-image";
const authorPostCardPreviewImg = "post-card-author-preview-image";
const ArticleHeader = "article-preview-title";
const ArticleSubtitle = "article-preview-subtitle";
const PostHeader = "post-preview-title";
const PostSubtitle = "post-preview-subtitle";
const PostAuthorName = "post-preview-author-name";
const PostPublishDate = "post-preview-publish-date";

const formSubmitAction = (event) => {
    event.preventDefault()
    console.log('отправка');
    getAllInputsData();
}
window.onload = () => {
    alert('страница загружена');
    let form = document.getElementById(formId);
    let formAuthorPhotoInput = document.getElementById(formAuthorAvatarId);
    let formHeroImg1Input = document.getElementById(formHeroImg1Id);
    let formHeroImg2Input = document.getElementById(formHeroImg2Id);
    let deleteAuthorAvatar = document.getElementById(removeAuthorAvatar);
    let deleteHeroImg1 = document.getElementById(removeHeroImg1);
    let deleteHeroImg2 = document.getElementById(removeHeroImg2);
    let postTitleInput = document.getElementById(formTitleId);
    let postSubtitleInput = document.getElementById(formSubtitleId);
    let postAuthorInput = document.getElementById(formAuthorNameId);
    let postDateInput = document.getElementById(formPublishDateId);

    if(form){
        form.addEventListener('submit', formSubmitAction);
    }

    if(formAuthorPhotoInput){
        formAuthorPhotoInput.addEventListener('change', () => {
            previewAndLoadPicture(
            formAuthorAvatarId, AuthorAvatarImage, defaultAuthorAvatarText, 
            AuthorAvatarImagePreviewDiv, AuthorAvatarSettings, authorAvatarData, authorPostCardPreviewImg)
        })
    }

    if(deleteAuthorAvatar){
        deleteAuthorAvatar.addEventListener('click', () => {
            deletePhoto( 
                formAuthorAvatarId, AuthorAvatarImage, defaultAuthorAvatarText, 
                AuthorAvatarImagePreviewDiv, AuthorAvatarSettings, authorPostCardPreviewImg)
        })
    }

    if(formHeroImg1Input){
        formHeroImg1Input.addEventListener('change', () => {
            previewAndLoadPicture(
                formHeroImg1Id, heroImg1Img, defaultHeroImg1Text, 
                heroImg1PreviewDiv, HeroImg1Settings, heroImg1Data, acticlePreviewImg, HeroImg1Description)
         })
    }

    if(deleteHeroImg1){
        deleteHeroImg1.addEventListener('click', () => {
            deletePhoto( 
                formHeroImg1Id, heroImg1Img, defaultHeroImg1Text, 
                heroImg1PreviewDiv, HeroImg1Settings, acticlePreviewImg, HeroImg1Description)
        })
    }

    if(formHeroImg2Input){
        formHeroImg2Input.addEventListener('change', () => {
            previewAndLoadPicture(
            formHeroImg2Id, heroImg2Img, defaultHeroImg2Text, 
            heroImg2PreviewDiv, HeroImg2Settings, heroImg2Data, postCardPreviewImg, HeroImg2Description)
         })
    }   

    if(deleteHeroImg2){
        deleteHeroImg2.addEventListener('click', () => {
            deletePhoto( 
                formHeroImg2Id, heroImg2Img, defaultHeroImg2Text, 
                heroImg2PreviewDiv, HeroImg2Settings, postCardPreviewImg, HeroImg2Description)
        })
    }

    if(postTitleInput){
        postTitleInput.addEventListener('change', () => {
            insertChangesIntoPreview(formTitleId, ArticleHeader);
            insertChangesIntoPreview(formTitleId, PostHeader);
        })
    }

    if(postSubtitleInput){
        postSubtitleInput.addEventListener('change', () => {
            insertChangesIntoPreview(formSubtitleId, ArticleSubtitle);
            insertChangesIntoPreview(formSubtitleId, PostSubtitle);
        })
    }
    if(postAuthorInput){
        postAuthorInput.addEventListener('change', () => {
            insertChangesIntoPreview(formAuthorNameId, PostAuthorName);
        })
    }

    if(postDateInput){
        postDateInput.addEventListener('change', () => {
            let publish = document.getElementById(PostPublishDate);
            let changes = getNormalDateValue(postDateInput.value);
            if(publish){
                publish.innerHTML = changes;
            }
        })
    }
}

const insertChangesIntoPreview = (changedItem, previewFieldId) => {
    let previewField = document.getElementById(previewFieldId);
    let changes = getInputsTextData(changedItem);
    if(previewField){
        previewField.innerHTML = changes;
    }
}

const deletePhoto = (
    pictureInputId, picturePreviewId, 
    inputDefaultTextId, picturePreviewDivId, 
    pictureSettingsDivId, previewAcrticlesImage, pictureInputDesctiption = "") => {

    let pictureInput = document.getElementById(pictureInputId);
    let previewPicture = document.getElementById(picturePreviewId);
    let inputDefaultText = document.getElementById(inputDefaultTextId);
    let imagePreviewDiv = document.getElementById(picturePreviewDivId);
    let settingsMenu = document.getElementById(pictureSettingsDivId);
    let pictureInputDesc = document.getElementById(pictureInputDesctiption);
    let previewAcrticleImage = document.getElementById(previewAcrticlesImage)

    if(pictureInput){
        pictureInput.value = null;
    }

    if(inputDefaultText){
        inputDefaultText.classList.remove("input-default-text-hidden");
    }

    if(previewPicture){
        previewPicture.src = "../static/img/admin/camera.svg"
        previewPicture.classList.remove("image-input");
    }

    if(imagePreviewDiv){
        imagePreviewDiv.classList.remove("unbordered-img");
    }

    if(settingsMenu){
        settingsMenu.classList.remove("input-img-settings_active");
    }
    
    if(pictureInputDesc){
        pictureInputDesc.classList.remove("input-default-text-hidden");
    }

    if(previewAcrticleImage){
        previewAcrticleImage.classList.remove("arcticle-preview-image_active");
        previewAcrticleImage.src = null;
    }
}

const previewAndLoadPicture = (
    pictureInputId, picturePreviewId, 
    inputDefaultText, picturePreviewDivId, 
    pictureSettingsDivId, pictureObj, previewAcrticlesImage, pictureInputDesctiption = "") => {

    let previewImg = document.getElementById(picturePreviewId);
    let previewDiv = document.getElementById(picturePreviewDivId);
    let textArea = document.getElementById(inputDefaultText);
    let pictureSettings = document.getElementById(pictureSettingsDivId);
    let imageFile = document.getElementById(pictureInputId).files[0];
    let pictureInputDesc = document.getElementById(pictureInputDesctiption);
    let previewArticleImg = document.getElementById(previewAcrticlesImage);
    let reader = new FileReader();

    reader.onloadend = () => {
        if(previewImg){
            previewImg.src = reader.result;
            pictureObj.base64 = reader.result;
        }
        if(previewArticleImg){
            previewArticleImg.classList.add("arcticle-preview-image_active");
            previewArticleImg.src = reader.result;
        }
    }

    if(previewDiv){
        previewDiv.classList.add("unbordered-img");
    }

    if(previewImg){
        previewImg.classList.add("image-input");
    }

    if(textArea){
        textArea.classList.add("input-default-text-hidden");
    }

    if(pictureSettings){
        pictureSettings.classList.add('input-img-settings_active');
    }

    if(pictureInputDesc){
        pictureInputDesc.classList.add("input-default-text-hidden");
    }

    let base64code;
    if(imageFile){
        reader.readAsDataURL(imageFile);
        reader.onerror = () => {
            console.log('there are some problems');
        }
    }
}

const getAllInputsData = () => {
    let inputsData = {};

    inputsData.title = getInputsTextData(formTitleId);
    inputsData.subtitle = getInputsTextData(formSubtitleId);
    inputsData.author = getInputsTextData(formAuthorNameId);
    inputsData.content = getInputsTextData(formContentId);
    inputsData.publish_date = getInputsTextData(formPublishDateId);
    inputsData.author_url = authorAvatarData.base64;
    inputsData.image_url = heroImg1Data.base64;
    inputsData.heroImg2 = heroImg2Data.base64;
    inputsData.publish_date = getNormalDateValue(inputsData.publish_date);
    let jsonInput = JSON.stringify(inputsData);
    console.log(jsonInput);
}

const getNormalDateValue = (date) => {
    let year = date.slice(0, date.indexOf('-', 0));
    let month = date.slice(date.indexOf('-', 0) + 1, date.indexOf('-', date.indexOf('-', 0) + 1));
    let day = date.slice(date.indexOf('-', date.indexOf('-', 0) + 1) + 1, date.length);
    return month + "/" + day + "/" + year 
}

const getInputsTextData = (inputId) => {
    let textInput = document.getElementById(inputId);
    if(textInput)
    {
        return textInput.value;
    } 
    else
    {
        return "";
    }
}

