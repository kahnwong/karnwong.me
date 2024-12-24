import os, glob
import frontmatter
from datetime import datetime
from icecream import ic
import shutil

os.makedirs("output", exist_ok=True)

files = glob.glob(
    "/Users/kahnwong/Git/kahnwong/fleet/karnwong.me/content/posts/20*/*/*.md"
)
for i in files:
    print(f"==== Processing: {i} ====")

    post = frontmatter.load(i)

    # --- out metadata ---
    date= datetime.strftime(post.metadata["date"], "%Y-%m-%d")

    year = datetime.strftime(post.metadata["date"], "%Y")
    month = datetime.strftime(post.metadata["date"], "%m")

    # --- content ---
    content = post.content

    # --- write to file: plain article version ---
    # slug = i.split("/")[-1].split(".")[0]

    # out_filename = f"output/{date}-{slug}.md"
    # out_content = frontmatter.dumps(frontmatter.Post(
    #     handler=frontmatter.default_handlers.TOMLHandler(),
    #     title=post.metadata["title"],
    #     date= date,
    #     path= f"/posts/{year}/{month}/{slug}",
    #     content=content,
    #     taxonomies={ "categories": [],"tags": [i.replace(' ', '-') for i in post.metadata["tags"]]},
    # ))
    # with open(out_filename, 'w') as f:
    #     f.write(out_content)
    #
    # os.remove(i)

    # --- write to file: article with images version ---
    slug = i.split("/")[-2].split(".")[0]
    out_dir = f"output/{date}-{slug}"
    os.makedirs(out_dir, exist_ok=True)
    out_content = frontmatter.dumps(frontmatter.Post(
        handler=frontmatter.default_handlers.TOMLHandler(),
        title=post.metadata["title"],
        date= date,
        path= f"/posts/{year}/{month}/{slug}",
        content=content,
        taxonomies={ "categories": [],"tags": [i.replace(' ', '-') for i in post.metadata["tags"]]},
    ))
    with open(f"{out_dir}/index.md", 'w') as f:
        f.write(out_content)

    og_image_dir = i.strip("index.md")+"images"
    shutil.move(og_image_dir, out_dir+"/")

    os.remove(i)

    break
