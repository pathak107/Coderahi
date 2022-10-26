import axios from 'axios'

const instance = axios.create({
    baseURL: 'http://localhost:8080/api/v1/',
    timeout: 1000,
    // headers: {'X-Custom-Header': 'foobar'}
});

export const getAllCourses = async () => {
    const res = await instance.get(`course?draft=true`)
    return res
}

export const getCourseByIDWithSectionsAndPosts = async (course_id) => {
    const res = await instance.get(`course/${course_id}?section=true&post=true&draft=true`)
    return res
}
export const createCourse = async ({ title, desc, cost }) => {
    const res = await instance.post('course', {
        title: title,
        desc_short: desc,
        cost: parseInt(cost)
    })
    return res
}
export const editCourse = async ({ title, desc, cost, markdown, html, course_id, categories, publish }) => {
    console.log(cost)
    const res = await instance.patch(`course/${course_id}`, {
        title,
        desc_short: desc,
        cost: parseInt(cost),
        markdown,
        html,
        categories,
        publish
    })
    return res
}
export const deleteCourse = async (course_id) => {
    const res = await instance.delete(`course/${course_id}`)
    return res
}

export const getPostByID = async (post_id) => {
    const res = await instance.get(`post/${post_id}`)
    return res
}

export const createPost = async ({ section_id, title, desc }) => {
    const res = await instance.post('post', {
        title: title,
        description: desc,
        section_id: parseInt(section_id)
    })
    return res
}

export const editPost = async ({ post_id, title, desc, publish }) => {
    const res = await instance.patch(`post/${post_id}`, {
        title: title,
        description: desc,
        publish
    })
    return res
}

export const editPostBody = async ({ markdown, html, post_id }) => {
    console.log(markdown)
    const res = await instance.patch(`post/${post_id}`, {
        markdown,
        html
    })
    return res
}

export const deletePost = async (post_id) => {
    const res = await instance.delete(`post/${post_id}`)
    return res
}

export const createSection = async ({ course_id, title, desc }) => {
    const res = await instance.post('course/section', {
        title: title,
        desc,
        course_id: parseInt(course_id)
    })
    return res
}
export const editSection = async ({ section_id, title, desc, publish }) => {
    const res = await instance.patch(`course/section/${section_id}`, {
        title: title,
        desc,
        publish
    })
    return res
}
export const deleteSection = async (section_id) => {
    const res = await instance.delete(`course/section/${section_id}`)
    return res
}

export const uploadCourseImage = async ({file, course_id})=>{
    let formData = new FormData()
    formData.append("file", file)
    const res = await instance.post(`course/upload/image/${course_id}`, formData, {
        headers:  {
            "Content-Type": "multipart/form-data",
          }
    })
    return res
}

export const updateOrderOfPost = async ({post_id, order}) =>{
    console.log(post_id)
    console.log(order)
    const res = await instance.patch(`post/change/order`, {
        post_id,
        order
    })
    return res
}

export const updateOrderOfSection = async ({section_id, order}) =>{
    const res = await instance.patch(`course/section/order`, {
        section_id,
        order
    })
    return res
}

export const getAllCategories = async ()=>{
    const res = await instance.get(`category`)
    return res
}