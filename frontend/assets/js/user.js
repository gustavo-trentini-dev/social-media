$('#unfollow').on("click", unfollow)
$('#follow').on("click", follow)
$('#edit-user').on("submit", edit)
$('#update-password').on("submit", updatePassword)
$('#delete-user').on("click", deleteUser)

function deleteUser() {
  Swal.fire({
    title: 'Are you sure?',
    text: 'This will delete your account forever!',
    showCancelButton: true,
    cancelButtonText: 'Cancel',
    icon: 'warning'
  }).then((confirmation) => {
    if (confirmation) {
      $.ajax({
        url: '/delete-user',
        method: 'DELETE'
      }).done(() => {
        Swal.fire("Success!", "Your user is deleted", "success")
        .then(() => {
          window.location = '/logout'
        })
      }).fail(() => {
        Swal.fire("Ops...", "Error on deleting account", "error")
      })
    }
  })

}

function edit(event) {
  event.preventDefault()
  $('#btn-edit').prop('disabled', true)

  const name = $('#name').val()
  const email = $('#email').val()
  const nick = $('#nick').val()

  $.ajax({
    url: `/edit-user`,
    method: "PUT",
    data: {
      name,
      email,
      nick
    }
  }).done(() => {
    Swal.fire('Success', 'User updated', 'success').then(() => {
      window.location = "/profile"
    })
  }).fail((err) => {
    Swal.fire('Ops...', 'Error on update user', 'error')
    $('#follow').prop('disabled', false)
  }).always(() => {
    $('#btn-edit').prop('disabled', false)
  })
}

function updatePassword(event) {
  event.preventDefault()

  const currentPassword = $('#current-password').val()
  const newPassword = $('#new-password').val()
  const confirmPassword = $('#confirm-password').val()

  if (newPassword != confirmPassword) {
    Swal.fire('Ops...', "New password don't match", 'warning')
  }

  $.ajax({
    url: `/update-password`,
    method: "PUT",
    data: {
      new: newPassword,
      current: currentPassword
    }
  }).done(() => {
    Swal.fire('Success', 'Password updated', 'success').then(() => {
      window.location = "/profile"
    })
  }).fail((err) => {
    Swal.fire('Ops...', 'Error on update password', 'error')
    $('#follow').prop('disabled', false)
  })
}

function follow(){
  const userId = $(this).data('user-id')
  $(this).prop('disabled', true)

  $.ajax({
    url: `/users/${userId}/follow`,
    method: "POST"
  }).done(() => {
    window.location.reload()
  }).fail((err) => {
    Swal.fire('Ops...', 'Error on follow user', 'error')
    $('#follow').prop('disabled', false)
  })
}

function unfollow() {
  const userId = $(this).data('user-id')
  $(this).prop('disabled', true)

  console.log('us ->', userId)

  $.ajax({
    url: `/users/${userId}/unfollow`,
    method: "POST"
  }).done(() => {
    window.location.reload()
  }).fail((err) => {
    Swal.fire('Ops...', 'Error on unfollow user', 'error')
    $('#unfollow').prop('disabled', false)
  })
}