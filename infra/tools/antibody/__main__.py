# Copyright 2015 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

"""Antibody is meant to audit reviews for the Chromium project.

Example invocation: [TBD]
  ./run.py infra.tools.antibody <arguments>
"""

# This file is untested, keep as little code as possible in there.

import argparse
import logging
import sys

from infra.tools.antibody import antibody
import infra.tools.antibody.cloudsql_connect as csql
from infra.tools.antibody import code_review_parse
from infra.tools.antibody import git_commit_parser
import infra_libs.logs


# https://storage.googleapis.com/chromium-infra-docs/infra/html/logging.html
LOGGER = logging.getLogger(__name__)


def main(argv):
  parser = argparse.ArgumentParser(
    prog="antibody",
    description=sys.modules['__main__'].__doc__)
  antibody.add_argparse_options(parser)
  infra_libs.logs.add_argparse_options(parser)
  args = parser.parse_args(argv)

  infra_libs.logs.process_argparse_options(args)

  # Do more processing here
  LOGGER.info('Antibody starting')
  with open(args.sql_password_file, 'r') as f:
    password = f.read()
  connection, cc = csql.connect(password)
  antibody.setup_antibody_db(cc)
  if args.rietveld_url:
    # TODO: get git hash from rietveld url
    code_review_parse.add_rietveld_data_to_db(None, args.rietveld_url, cc)
  else:
    git_commit_parser.upload_git_to_sql(cc)
    git_commits_with_review_urls = git_commit_parser.get_urls_from_git_db(cc)
    for git_hash, review_url in git_commits_with_review_urls:
      # cannot get access into chromereview.googleplex.com
      if 'chromereviews.googleplex' not in review_url:
        code_review_parse.add_rietveld_data_to_db(git_hash, review_url, cc)
  print code_review_parse.get_tbr_no_lgtm(cc)
  csql.close(connection, cc)


if __name__ == '__main__':
  sys.exit(main(sys.argv[1:]))