$('#new-post').on('submit', createPost);
$('#update-post').on('click', updatePost);
$('.delete-post').on('click', deletePost);

$(document).on('click', '.like-post', likePost)
$(document).on('click', '.unlike-post', unlikePost)

function createPost(event) {
  event.preventDefault();
  const title = $('#title').val()
  const content = $('#content').val()

  $.ajax({
    url: '/create-post',
    method: "POST",
    data: {
      title,
      content
    }
  }).done(() => {
    window.location = '/home'
  }).fail((err) => {
    Swal.fire('Ops...', 'Error on create post', 'error')
  });
}

function updatePost() {
  $(this).prop('disabled', true)
  const postId = $(this).data('post-id')
  const title = $('#title').val()
  const content = $('#content').val()

  $.ajax({
    url: `/update-post/${postId}`,
    method: "PUT",
    data: {
      title,
      content
    }
  }).done(() => {
    Swal.fire(
      "Done!",
      "Post updated...",
      "success"
    ).then(() => {
      window.location = "/home"
    })
  }).fail(() => {
    Swal.fire('Ops...', 'Error on update post', 'error')
  }).always(() => {
    $('#update-post').prop('disabled', false)
  })
}

function deletePost(event) {
  event.preventDefault()

  Swal.fire({
    title: 'Hey friend!',
    text: 'Are you sure you wanna delete this post?',
    showCancelButton: true,
    cancelButtonText: "Cancel",
    icon: "warning"
  }).then((confirmation) => {
    if(!confirmation.value) return

    const element = $(event.target)
    const postDiv = element.closest('div')
    const postId = postDiv.data('post-id')

    element.prop('disabled', true)

    $.ajax({
      url: `/delete-post/${postId}`,
      method: "DELETE"
    }).done(() => {
      postDiv.fadeOut("slow", () => {
        $(this).remove()
      })
    }).fail((err) => {
      Swal.fire('Ops...', 'Error on delete post', 'error')
    }).always(() => {
      element.prop('disabled', false)
    })
  })  
}

function likePost(event) {
  event.preventDefault()
  const element = $(event.target)
  const postId = element.closest('div').data('post-id')

  element.prop('disabled', true)

  $.ajax({
    url: `/posts/${postId}/like`,
    method: "POST"
  }).done(() => {
    const countLikes = element.next('span')
    const likes = parseInt(countLikes.text())
    
    countLikes.text(likes + 1)

    element.addClass('unlike-post')
    element.addClass('text-danger')
    element.removeClass('like-post')
  }).fail((err) => {
    Swal.fire('Ops...', 'Error on like post', 'error')
  }).always(() => {
    element.prop('disabled', false)
  })
}

function unlikePost(event) {
  event.preventDefault()
  const element = $(event.target)
  const postId = element.closest('div').data('post-id')

  element.prop('disabled', true)

  $.ajax({
    url: `/posts/${postId}/dislike`,
    method: "POST"
  }).done(() => {
    const countLikes = element.next('span')
    const likes = parseInt(countLikes.text())
    
    countLikes.text(likes - 1)

    element.addClass('like-post')
    element.removeClass('unlike-post')
    element.removeClass('text-danger')
  }).fail((err) => {
    Swal.fire('Ops...', 'Error on dislike post', 'error')
  }).always(() => {
    element.prop('disabled', false)
  })
}