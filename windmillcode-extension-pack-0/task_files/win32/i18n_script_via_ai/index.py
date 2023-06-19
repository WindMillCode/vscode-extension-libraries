import argparse
import json
import os
import re
import time

def local_deps():
  import sys
  if sys.platform == 'win32':
    sys.path.append(sys.path[0] + '\site-packages\windows')
  elif sys.platform =='linux':
    sys.path.append(sys.path[0] + '/site-packages/linux')
  elif sys.platform =='darwin':
    sys.path.append(sys.path[0] + '/site-packages/linux')
local_deps()
import openai
import pprint
pp = pprint.PrettyPrinter(indent=2, compact=False, width=1)


def print_if_dev(item,pretty=False):
        if pretty == True:
          pp.pprint(item)
        else:
          print(item)


class OpenAIManager():
  init= False
  client = None
  def __init__(self,api_key):
      self.init = True
      openai.api_key = api_key

  def _ask_chatgpt(self,prompt,randomness=0):
    try:
      response = openai.ChatCompletion.create(
          model="gpt-3.5-turbo-0301",
          messages=[{
            "role":"user",
            "content":prompt,
          }],
          temperature=randomness,
      )
    except BaseException as e:
      time.sleep(21)
      return self._ask_chatgpt(prompt,randomness)

    # response = openai.Completion.create(
    #   model="text-davinci-003",
    #   prompt=prompt,
    #   temperature=randomness
    # )
    print_if_dev(response,True)

    return response.choices[0].message.content


  def i18n_overwrite_translations(self,dev_obj):


    lang_codes = dev_obj.get("lang_codes")
    source_file = dev_obj.get("source_file")
    dest_file = dev_obj.get("dest_file")
    abs_path_source_file = dev_obj.get("abs_path_source_file")
    source_language=source_file.split(".")[0]
    for x in lang_codes:

        with open(abs_path_source_file,encoding="utf-8") as f:
            lang  = json.load(f)
            new_lang ={}
            for k,v in lang.items():
              print(k)
              if len(v.items()) == 0:
                new_lang[k] = {}
              else:
                prompt = """ With the following JSON {} recursively translate every value from English  to the  {} language  `{}`
                make sure to return proper JSON and use the following text for mapping reference
                zh stands for Mandarin Chinese,
                es stands for Spanish,
                hi stands for Hindi,
                uk stands for Ukranian,
                ar stands for Arabic,
                bn stands for Bengali,
                ms stands for Malay,
                fr stands for French,
                de stands for German,
                sw stands for Swahili,
                and am stands for Amharic
                """.format(
                  v,source_language,x,
                )
              my_translate = self._ask_chatgpt(prompt)
              my_translate = re.sub(r'[\\\n]', '', my_translate)
              try:
                new_lang[k] = json.loads(my_translate)
              except json.JSONDecodeError:
                new_lang[k] ={}

            abs_path_dest_file = os.path.join(os.getcwd(),args.location,dest_file.replace("{}",x))
            with open(abs_path_dest_file,"w",encoding="utf-8") as g:
                print(json.dumps(new_lang,indent=2) , file=g)
                f.close()
                g.close()

  def fix_translations(self,dev_obj):


    lang_codes = dev_obj.get("lang_codes")
    source_file = dev_obj.get("source_file")
    dest_file = dev_obj.get("dest_file")
    abs_path_source_file = dev_obj.get("abs_path_source_file")
    source_language=source_file.split(".")[0]
    for x in lang_codes:

      abs_path_dest_file = os.path.join(os.getcwd(),args.location,dest_file.replace("{}",x))
      with open(abs_path_dest_file,"r",encoding="utf-8") as g:
        lang_string =g.read()

        lang  = json.loads(lang_string)
        new_lang ={}
        for k,v in lang.items():
          print(k)

          try:
            new_lang[k] = json.loads(v)
          except BaseException as e:
            new_lang[k] = {}
          # print(json.dumps(new_lang,indent=2) , file=g)
        with open(abs_path_dest_file,"w",encoding="utf-8") as h:
          json.dump(new_lang,indent=2,fp=h)
          g.close()
          h.close()


if __name__ == "__main__":
    parser = argparse.ArgumentParser(
                        prog='Translation Scirpt',
                        description='translates angular 18n script',
                        epilog='Text at the bottom of help')
    parser.add_argument('-l','--location')
    parser.add_argument('-s','--source-file')
    parser.add_argument('-d','--dest-file',default="{}.json")
    parser.add_argument('-c','--lang-codes')
    args = parser.parse_args()
    abs_path_source_file = os.path.join(os.getcwd(),args.location,args.source_file)

    lang_codes = args.lang_codes.split(",")
    params= {
        "lang_codes":lang_codes,
        "source_file":args.source_file,
        "dest_file":args.dest_file,
        "abs_path_source_file":abs_path_source_file
    }
    mngr = OpenAIManager(os.environ.get("OPENAI_API_KEY_0"))
    mngr.i18n_overwrite_translations(params)
    # mngr.fix_translations(params)
