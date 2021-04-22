$('#follow').on('click', follow);
$('#unfollow').on('click', unfollow);
$('#delete-user').on('click', deleteUser);

$('#edit-user').on('submit', update);
$('#update-password').on('submit', updatePassword);

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

function updatePassword(event) {
  event.preventDefault();

  if ($('#new-password').val() != $('#confirmation-password').val()) {
    Swal.fire(
      'Oops...',
      "Your new password doesn't match the confirmation field!",
      'warning'
    );
    return;
  }

  $.ajax({
    url: "/update-password",
    method: "POST",
    data: {
      current: $('#current-password').val(),
      new: $('#new-password').val()
    }
  }).done(function () {
    Swal.fire(
      "Alright!",
      "Your password has been updated!",
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

function deleteUser() {
  Swal.fire({
    title: "Warning!",
    text: "Are you sure you want to delete your account? This cannot be undone.",
    showCancelButton: true,
    cancelButtonText: "Cancel",
    icon: "warning"
  }).then(function (confirmation) {
    if (confirmation.value) {
      $.ajax({
        url: "/delete-account",
        method: "DELETE"
      }).done(function () {
        Swal.fire(
          "Alright!",
          "Your account has been deleted successfuly!",
          "success"
        ).then(function () {
          window.location = "/logout";
        })
      }).fail(function () {
        Swal.fire(
          'Oops...',
          "We couldn't delete your account at the moment. Please try again later.",
          'error'
        );
      })
    }
  })
}