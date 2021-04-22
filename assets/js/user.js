$('#follow').on('click', follow);
$('#unfollow').on('click', unfollow);
$('#edit-user').on('submit', update);

function follow() {
  const userID = $(this).data('user-id');
  $(this).prop('disabled', true);

  $.ajax({
    url: `/users/${userID}/follow`,
    method: "POST"
  }).done(function () {
    window.location = `/users/${userID}`;
  }).fail(function () {
    Swal.fire(
      'Oops...',
      "Something went wrong. Please try again later.",
      'error'
    );
    $(this).prop('disabled', false);
  });
}

function unfollow() {
  const userID = $(this).data('user-id');
  $(this).prop('disabled', true);

  $.ajax({
    url: `/users/${userID}/unfollow`,
    method: "DELETE"
  }).done(function () {
    window.location = `/users/${userID}`;
  }).fail(function () {
    Swal.fire(
      'Oops...',
      "Something went wrong. Please try again later.",
      'error'
    );
    $(this).prop('disabled', false);
  });
}

function update(event) {
  event.preventDefault();

  $.ajax({
    url: "update-user",
    method: "PUT",
    data: {
      name: $('#name').val(),
      email: $('#email').val(),
      username: $('#username').val(),
    }
  }).done(function () {
    Swal.fire(
      "Alright!",
      "You've successfuly updated your personal information",
      "success"
    ).then(function () {
      window.location = "/profile";
    })
  }).fail(function () {
    Swal.fire(
      'Oops...',
      "Something went wrong. Please try again later.",
      'error'
    );
  })
}