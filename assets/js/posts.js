$('#new-post').on('submit', createPost);
$(document).on('click', '.like-post', likePost)
$(document).on('click', '.dislike-post', dislikePost)

function createPost(event) {
  event.preventDefault();

  $.ajax({
    url: "/posts",
    method: "POST",
    data: {
      title: $('#title').val(),
      content: $('#content').val(),
    }
  }).done(function () {
    window.location = "/home";
  }).fail(function () {
    alert("Failed to submit post. Please try again.");
  })
}

function likePost(event) {
  event.preventDefault();

  const clickedElement = $(event.target);
  const postID = clickedElement.closest('div').data('post-id');

  clickedElement.prop('disabled', true);

  $.ajax({
    url: `/posts/${postID}/like`,
    method: "POST"
  }).done(function () {

    const likesCounter = clickedElement.next('span');
    const likes = parseInt(likesCounter.text());
    var likesText = likes + 1

    if (likesText == 1) {
      likesText = likesText + " like";
    } else {
      likesText = likesText + " likes";
    }

    likesCounter.text(likesText)

    clickedElement[0].className = "fas fa-heart dislike-post text-danger";

  }).fail(function () {
    alert("Something went wrong")
  }).always(function () {
    clickedElement.prop('disabled', false);
  });
}

function dislikePost(event) {
  event.preventDefault();

  const clickedElement = $(event.target);
  const postID = clickedElement.closest('div').data('post-id');

  clickedElement.prop('disabled', true);

  $.ajax({
    url: `/posts/${postID}/dislike`,
    method: "POST"
  }).done(function () {

    const likesCounter = clickedElement.next('span');
    const likes = parseInt(likesCounter.text());
    var likesText = likes - 1

    if (likesText == 1) {
      likesText = likesText + " like";
    } else {
      likesText = likesText + " likes";
    }

    likesCounter.text(likesText)

    clickedElement[0].className = "far fa-heart like-post";

  }).fail(function () {
    alert("Something went wrong")
  }).always(function () {
    clickedElement.prop('disabled', false);
  });
}