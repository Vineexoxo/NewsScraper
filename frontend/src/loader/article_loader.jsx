import { Request } from "@/request/request";

export const ArticleLoader = async ({params}) => {
    const {url}=params;
    console.log("url", url);

    let data = await Request.post('/api/v1/articles/get/', {filters:[{field:"Link", value:url, comparison:"eq"}]});
    console.log("Dataigger loader", data);
    return data.data;
}