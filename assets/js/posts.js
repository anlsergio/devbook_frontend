$('#new-post').on('submit', createPost);

$(document).on('click', '.like-post', likePost)
$(document).on('click', '.dislike-post', dislikePost)

$('#update-post').on('click', updatePost);
$('.delete-post').on('click', deletePost);

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
    Swal.fire(
      'Yay!',
      'Your new post has been submitted!',
      'success'
    ).then(function () {
      window.location = "/home";
    })
  }).fail(function () {
    Swal.fire(
      'Oops...',
      "Failed to submit post. Please try again later.",
      'error'
    )
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
    Swal.fire(
      'Oops...',
      "Something went wrong. Please try again later.",
      'error'
    )
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
    Swal.fire(
      'Oops...',
      "Something went wrong. Please try again later.",
      'error'
    )
  }).always(function () {
    clickedElement.prop('disabled', false);
  });
}

function updatePost() {
  $(this).prop('disabled', true);

  const postID = $(this).data('post-id');

  $.ajax({
    url: `/posts/${postID}`,
    method: "PUT",
    data: {
      title: $('#title').val(),
      content: $('#content').val()
    }
  }).done(function () {
    Swal.fire(
      'Done!',
      'Your post has been updated!',
      'success'
    ).then(function () {
      window.location = "/home";
    })
  }).fail(function () {
    Swal.fire(
      'Oops...',
      "Something went wrong and your post couldn't be updated. Please try again later.",
      'error'
    )
  }).always(function () {
    $('#update-post').prop('disabled', false);
  });
}

function deletePost(event) {
  event.preventDefault();

  Swal.fire({
    title: 'Confirmation',
    text: 'Are you sure you want to delete this post?',
    showCancelButton: true,
    cancelButtonText: 'No, take me back!',
    icon: 'warning'
  }).then(function (confirmation) {
    if (!confirmation.value) return;

    const clickedElement = $(event.target);
    const post = clickedElement.closest('div');
    const postID = post.data('post-id');

    clickedElement.prop('disabled', true);

    $.ajax({
      url: `/posts/${postID}`,
      method: "DELETE"
    }).done(function () {
      post.fadeOut("slow", function () {
        $(this).remove();
      });
    }).fail(function () {
      Swal.fire(
        'Oops...',
        "Something went wrong while trying to delete the post. Please try again later.",
        'error'
      )
    });
  })

}