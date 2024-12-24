import os, glob
import frontmatter
from datetime import datetime
from icecream import ic

os.makedirs("output", exist_ok=True)

files = glob.glob(
    "/Users/kahnwong/Git/kahnwong/fleet/karnwong.me/content/posts/2023/*.md"
)
for i in files:
    print(f"==== Processing: {i} ====")

    post = frontmatter.load(i)

    # --- out metadata ---
    date= datetime.strftime(post.metadata["date"], "%Y-%m-%d")

    year = datetime.strftime(post.metadata["date"], "%Y")
    month = datetime.strftime(post.metadata["date"], "%m")

    slug = i.split("/")[-1].split(".")[0]

    # --- content ---
    content = post.content

    # --- write to file ---
    out_filename = f"output/{date}-{slug}.md"
    out_content = frontmatter.dumps(frontmatter.Post(
        handler=frontmatter.default_handlers.TOMLHandler(),
        title=post.metadata["title"],
        date= date,
        path= f"/posts/{year}/{month}/{slug}",
        content=content,
        taxonomies={ "categories": [],"tags": [i.replace(' ', '-') for i in post.metadata["tags"]]},
    ))
    with open(out_filename, 'w') as f:
        f.write(out_content)

    os.remove(i)

    # break
