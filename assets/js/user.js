$('#follow').on('click', follow);
$('#unfollow').on('click', unfollow);

function follow() {
  const userID = $(this).data('user-id');
  $(this).prop('disabled', true);

  $.ajax({
    url: `/users/${userID}/follow`,
    method: "POST"
  }).done(function() {
    window.location = `/users/${userID}`;
  }).fail(function() {
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
  }).done(function() {
    window.location = `/users/${userID}`;
  }).fail(function() {
    Swal.fire(
      'Oops...',
      "Something went wrong. Please try again later.",
      'error'
    );
    $(this).prop('disabled', false);
  });
}